package com.demo.service;


import com.demo.entity.User;

public interface IUserService {
    //根据账号密码获取用户
    User getUserByUsernameAndPassword(String username,String password);
    //根据账号获取用户
    User getUserByUserName(String username);
    //新增用户
    Integer addUser(String username,String password);
    //更新用户
    Integer updateUser(User user);

}
