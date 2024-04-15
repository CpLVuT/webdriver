package com.demo.controller;

import com.demo.entity.File;
import com.demo.entity.User;
import com.demo.service.IFileService;
import com.demo.service.IUserService;
import com.demo.utils.Authentication;
import org.json.JSONException;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.Part;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.sql.*;
import java.util.*;
import java.util.zip.ZipEntry;
import java.util.zip.ZipOutputStream;

@RestController
@RequestMapping("/file")
public class FileController {
    @Autowired
    private IUserService userService;
    @Autowired
    private IFileService fileService;

    @GetMapping("/listPage")
    public void listPage(int page,int itemsPerPage,String searchFileName,String searchFileStartTime,String searchFileEndTime,HttpServletRequest request, HttpServletResponse response) throws IOException {
        response.setContentType("text/html; charset=utf-8");
        if (!Authentication.isLogin(request)) {
            return;
        }
        String userName = (String)request.getSession().getAttribute("username");
        try {
            User user = userService.getUserByUserName(userName);
            //计算总页数
            Integer fileCount = fileService.countByUserId(String.valueOf(user.getId()));
            int totalPages = fileCount/itemsPerPage + 1;
            //获取文件列表
            List<File> list = fileService.listPage(page,itemsPerPage,searchFileName,searchFileStartTime,searchFileEndTime,user.getId());

            JSONObject data = new JSONObject();
            data.put("totalPages",totalPages);
            data.put("files",list);

            //返回数据
            response.getWriter().print(data);

        } catch (JSONException e) {
            e.printStackTrace();
        }
    }

    @PostMapping("/upload")
    public void upload(HttpServletRequest request, HttpServletResponse response) throws IOException, ServletException {
        if(!Authentication.isLogin(request)){
            return;
        }
        String userName = (String)request.getSession().getAttribute("username");
        response.setContentType("text/html; charset=utf-8");
        Timestamp updated_at = new Timestamp(System.currentTimeMillis());
        System.out.println(updated_at);

        //获取文件对象
        Part part=request.getPart("file");

        String filename=part.getSubmittedFileName();

        //文件上传后默认路径在D盘的mydisk下面 确保文件夹存在
        String path="D://mydisk/"+filename;


        //保存
        try{
            part.write(path);
        }catch (Exception e){
            JSONObject jsonObject = new JSONObject();
            jsonObject.put("msg","上传路径默认在D盘下的mydisk文件，请确保文件目录存在！");
            response.setStatus(HttpServletResponse.SC_INTERNAL_SERVER_ERROR);
            response.getWriter().print(jsonObject);
            return;
        }

        User user = userService.getUserByUserName(userName);
        //判断数据库中是否有名字和路径一致的文件，若有则更新时间，若没有则添加新纪录
        File file = fileService.getFileByNameAndPath(user.getId(),filename,path);
        if (file != null){
            fileService.updateFile(user.getId(),filename,path, Math.toIntExact(part.getSize()));
        }else {
            fileService.addFile(user.getId(),filename,path, Math.toIntExact(part.getSize()));
        }

        response.sendRedirect("../disc.html");
    }

    @GetMapping("/moveFile")
    public void moveFile(String fileIds,String movePath,HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        if(!Authentication.isLogin(request)) {
            return;
        }
        List<String> fileIdList = Arrays.asList(fileIds.split(","));
        fileIdList.forEach(e -> {
            File file = fileService.loadById(Integer.valueOf(e));
            // 确认源文件存在
            if (!Files.exists(Paths.get(file.getPath()))) {
                try {
                    throw new Exception("Source file does not exist: " + file.getPath());
                } catch (Exception ex) {
                    throw new RuntimeException(ex);
                }
            }

            // 确保目标文件所在的目录存在
            if (!Files.exists(Paths.get(movePath))) {
                try {
                    Files.createDirectories(Paths.get(movePath)); // 创建所有不存在的父目录
                } catch (IOException ex) {
                    throw new RuntimeException(ex);
                }
            }
            // 计算移动后的文件路径
            Path targetFilePath = Paths.get(movePath).resolve(Paths.get(file.getPath()).getFileName());
            // 执行文件移动操作
            try {
                Files.move(Paths.get(file.getPath()), targetFilePath, StandardCopyOption.REPLACE_EXISTING);
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
            //移动完文件执行数据库记录更改
            file.setPath(movePath + "/" +  file.getFilename());
            fileService.update(file);
        });
    }

    @GetMapping("/downloadFile")
    public void downloadFile(String fileIds,HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

        if(!Authentication.isLogin(request)){
            return;
        }

        if (fileIds == null) {
            System.out.println("无文件ID参数");
            return;
        }
        List<String> fileIdList = Arrays.asList(fileIds.split(","));
        // 创建ZipOutputStream来写入响应的OutputStream中
        ZipOutputStream zipOut = new ZipOutputStream(response.getOutputStream());
        fileIdList.forEach(e -> {
            File file = fileService.loadById(Integer.valueOf(e));
            java.io.File fileToZip = new java.io.File(file.getPath());
            if(!fileToZip.exists()){
                System.out.println("路径'" + file.getPath() + "'下的文件'" + file.getFilename() + "'不存在");
                return;
            }
            // 创建ZipEntry，并添加到zip中
            try {
                zipOut.putNextEntry(new ZipEntry(fileToZip.getName()));
                Files.copy(fileToZip.toPath(), zipOut);
                zipOut.closeEntry();
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }

        });

        // 完成zip文件的创建
        zipOut.finish();
        zipOut.close();
        response.setContentType("application/zip");
        response.setHeader("Content-Disposition", "attachment; filename=\"download.zip\"");

    }

    @GetMapping("/deleteFiles")
    public void deleteFiles(String fileIds,HttpServletRequest request, HttpServletResponse response) {
        if (!Authentication.isLogin(request)) {
            return;
        }
        List<String> fileIdList = Arrays.asList(fileIds.split(","));
        fileIdList.forEach(e -> {
            File result = fileService.loadById(Integer.valueOf(e));
            java.io.File file = new java.io.File(result.getPath());
            if (file.exists()) {
                boolean isDeleted = file.delete();
                if (isDeleted) {
                    System.out.println("文件删除成功 文件名： " + result.getFilename() + "  文件路径：" + result.getFilename());
                } else {
                    System.out.println("文件删除失败 文件名： " + result.getFilename() + "  文件路径：" + result.getFilename());
                    // 发送失败响应
                    response.setStatus(HttpServletResponse.SC_INTERNAL_SERVER_ERROR);
                    try {
                        response.getWriter().write("文件删除失败 文件名： " + result.getFilename() + "  文件路径：" + result.getFilename());
                    } catch (IOException ex) {
                        throw new RuntimeException(ex);
                    }
                    return;
                }
            }
            fileService.deleteFileById(e);
        });

        System.out.println(fileIds);
    }

}
