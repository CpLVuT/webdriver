package com.demo.service.Impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.toolkit.StringUtils;
import com.demo.entity.File;
import com.demo.mapper.FileMapper;
import com.demo.service.IFileService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.math.BigInteger;
import java.util.Date;
import java.util.List;
@Service
public class IFileServiceImpl implements IFileService {
    @Resource
    private FileMapper mapper;
    @Override
    public List<File> listPage(int page, int itemsPerPage, String searchFileName, String searchFileStartTime, String searchFileEndTime,Integer userId) {
        LambdaQueryWrapper<File> queryWrapper = new LambdaQueryWrapper();
        queryWrapper.like(StringUtils.isNotBlank(searchFileName),File::getFilename,searchFileName)
                .gt(StringUtils.isNotBlank(searchFileStartTime),File::getTime,searchFileStartTime)
                .lt(StringUtils.isNotBlank(searchFileEndTime),File::getTime,searchFileEndTime)
                .eq(File::getUserId,userId)
                .last("limit " + itemsPerPage + " OFFSET " + (page -1) *itemsPerPage);
        List<File> list = mapper.selectList(queryWrapper);

        return list;
    }

    @Override
    public Integer countByUserId(String userId) {
        LambdaQueryWrapper<File> queryWrapper = new LambdaQueryWrapper();
        queryWrapper.eq(File::getUserId,userId);
        Integer count = mapper.selectCount(queryWrapper);
        return count;
    }

    @Override
    public Integer addFile(Integer userId, String fileName, String path, Integer fileSize) {
        File file = new File();
        file.setUserId(userId);
        file.setFilename(fileName);
        file.setPath(path);
        file.setFilesize(BigInteger.valueOf(fileSize));
        file.setTime(new Date());
        return mapper.insert(file);
    }

    @Override
    public File getFileByNameAndPath(Integer userId, String fileName, String path) {
        LambdaQueryWrapper<File> queryWrapper = new LambdaQueryWrapper<>();
        queryWrapper.eq(File::getUserId,userId)
                .eq(File::getFilename,fileName)
                .eq(File::getPath,path);
        return mapper.selectOne(queryWrapper);
    }

    @Override
    public Integer updateFile(Integer userId, String fileName, String path, Integer fileSize) {
        File file = new File();
        file.setUserId(userId);
        file.setFilename(fileName);
        file.setPath(path);
        file.setFilesize(BigInteger.valueOf(fileSize));
        file.setTime(new Date());
        LambdaQueryWrapper<File> queryWrapper = new LambdaQueryWrapper<>();
        queryWrapper.eq(File::getUserId,userId)
                .eq(File::getFilename,fileName)
                .eq(File::getPath,path);
        return mapper.update(file,queryWrapper);

    }

    @Override
    public Integer deleteFile(Integer userId, String fileName, String path) {
        LambdaQueryWrapper<File> queryWrapper = new LambdaQueryWrapper<>();
        queryWrapper.eq(File::getUserId,userId)
                .eq(File::getFilename,fileName)
                .eq(File::getPath,path);
        return mapper.delete(queryWrapper);
    }

    @Override
    public File loadById(Integer id) {
        return mapper.selectById(id);
    }

    @Override
    public Integer update(File file) {
        file.setTime(new Date());
        return mapper.updateById(file);
    }

    @Override
    public Integer deleteFileById(String id) {
        return mapper.deleteById(id);
    }


}
