package com.demo.controller;

import com.demo.entity.User;
import com.demo.service.IUserService;
import com.demo.utils.Authentication;
import com.demo.utils.Encoder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    private IUserService userService;

    @PostMapping("/userLogin")
    public void userLogin(String username,String password,HttpServletRequest request,HttpServletResponse response) throws IOException {
        //密码加密
        password = Encoder.encodeBase64(password);
        User user = userService.getUserByUsernameAndPassword(username,password);
        if (user == null){
            System.out.println("用户名或密码错误");
            response.sendRedirect("login.html?error=true");
        }else{
            //保存登录状态
            request.getSession().setAttribute("username", username);
            response.sendRedirect("../disc.html");
        }
    }

    @PostMapping("/userRegister")
    public void userRegister(String username,String password,HttpServletResponse response) throws IOException {
        //查找是否有重复用户，有则报错
        User result = userService.getUserByUserName(username);
        if (result != null){
            System.out.println("用户已存在");
            response.sendRedirect("register.html?error=true");
            return;
        }
        //加密密码
        password = Encoder.encodeBase64(password);
        //将注册用户的信息插入数据库
        userService.addUser(username,password);
        response.sendRedirect("../user/login.html");

    }

    @PostMapping("/userResetPwd")
    public void userResetPwd(String username,String password, HttpServletResponse response) throws ServletException, IOException {
        //查找是否有用户，没有则报错
        User result = userService.getUserByUserName(username);
        if (result == null){
            System.out.println(username + " 用户不存在");
            response.sendRedirect("register.html?error=true");
            return;
        }
        //加密密码
        password = Encoder.encodeBase64(password);
        result.setPassword(password);
        userService.updateUser(result);
        response.sendRedirect("../user/login.html");
    }

    @GetMapping("/delUsername")
    public void delUsername(HttpServletRequest request){

        if(!Authentication.isLogin(request)){
            return;
        }
        String usernamePara=request.getParameter("username");

        String sessionUsername=(String)request.getSession().getAttribute("username");

        if(usernamePara.equals(sessionUsername)){
            request.getSession().invalidate();
            System.out.println("用户"+usernamePara+"退出登录");
        }
    }

    @GetMapping("/getUsername")
    public void getUsername(HttpServletRequest request, HttpServletResponse response) throws  IOException {
        response.setContentType("text/html;charset=utf-8");
        response.getWriter().print(request.getSession().getAttribute("username"));
    }

}
