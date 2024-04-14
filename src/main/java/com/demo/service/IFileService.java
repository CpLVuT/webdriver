package com.demo.service;


import com.demo.entity.File;

import java.util.List;

public interface IFileService {
    //分页获取
    List<File> listPage(int page,int itemsPerPage,String searchFileName,String searchFileStartTime,String searchFileEndTime,Integer userId);
    //根据用户id获取文件数量
    Integer countByUserId(String userId);
    //新增
    Integer addFile(Integer userId,String fileName,String path,Integer fileSize);
    //查找
    File getFileByNameAndPath(Integer userId,String fileName,String path);
    //更新
    Integer updateFile(Integer userId,String fileName,String path,Integer fileSize);
    //删除
    Integer deleteFile(Integer userId,String fileName,String path);
    //id查找
    File loadById(Integer id);
    //全更新
    Integer update(File file);
    //id删除
    Integer deleteFileById(String id);
}
