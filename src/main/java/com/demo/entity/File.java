package com.demo.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigInteger;
import java.util.Date;

@TableName("file")
@Data
@AllArgsConstructor
@NoArgsConstructor
public class File {
    @TableId(type = IdType.AUTO)
    private Integer id;
    private Integer userId;
    private String filename;
    private BigInteger filesize;
    private Date time;
    private String path;
}
