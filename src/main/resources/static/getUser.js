let ajax=new XMLHttpRequest();
ajax.onreadystatechange=function(){
    if(ajax.readyState==4 && ajax.status==200){
        //做什么
        document.getElementById("username").innerText=ajax.responseText;
      //  document.cookie="username="+ajax.responseText;

//         //拉取文件列表
//         let ajaxfl=new XMLHttpRequest();
//         ajaxfl.onreadystatechange=function(){
//             if(ajaxfl.readyState==4 && ajaxfl.status==200){
//                 //做什么
//                 let data=JSON.parse(ajaxfl.responseText);
//                 let len=data.files.length;
//                 let table=document.getElementById("fileList");
//                 for(let i=0;i<len;i++){
//
//                     let selectorTd=document.createElement("td");
//                     let selector=document.createElement("input");
//                     selector.setAttribute("type","radio");
//                     selector.setAttribute("name","filename");
//                     selector.setAttribute("value",data.files[i].filename);
//                     selectorTd.appendChild(selector);
//
//                     let name=document.createElement("td");
//                     name.appendChild(document.createTextNode(data.files[i].filename));
//
//                     let size=document.createElement("td");
//                     size.appendChild(document.createTextNode(data.files[i].filesize));
//
//                     let time=document.createElement("td");
//                     time.appendChild(document.createTextNode(data.files[i].time));
//
//                     let row=document.createElement("tr");
//
//                     row.appendChild(selectorTd);
//                     row.appendChild(name);
//                     row.appendChild(size);
//                     row.appendChild(time);
//                     table.appendChild(row);
//                 }
//             }
//         };
//
// //发给谁
//         ajaxfl.open("post","getFileList",true);
//         ajaxfl.send();

    }
};

//发给谁
ajax.open("get","/user/getUsername",true);
ajax.send();


function logout() {

    //找到id=username 的标签 ，获取他的内部文本
    let username= document.getElementById("username").innerText;
    let ajax2=new XMLHttpRequest();

    ajax2.onreadystatechange=function(){
        if(ajax2.readyState==4 && ajax2.status==200){
        }
    };
    //发给谁
    ajax2.open("get","/user/delUsername?username="+username,true);
    ajax2.send();
    window.location.href="user/login.html"
}



let currentPage = 1;
const itemsPerPage = 10;
let totalPages = 1; // 假设有10页数据，这个值应该从后端获取

function fetchFiles(page) {
    var xhr = new XMLHttpRequest();
    let searchFileName = document.getElementById("searchFileName").value;
    let searchFileStartTime = document.getElementById("searchFileStartTime").value;
    let searchFileEndTime = document.getElementById("searchFileEndTime").value;
    if(searchFileStartTime.length > 0){
        searchFileStartTime = searchFileStartTime.replace('T',' ');
    }
    if(searchFileEndTime.length > 0){
        searchFileEndTime = searchFileEndTime.replace('T',' ');
    }
    let url = "/file/listPage?page=" + page + "&itemsPerPage=" + itemsPerPage + "&searchFileName=" + searchFileName
                        + "&searchFileStartTime=" + searchFileStartTime + "&searchFileEndTime=" + searchFileEndTime;
    xhr.open("GET",url, true);
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            var response = JSON.parse(xhr.responseText);
            var files = response.files;
            let filesList = document.getElementById('filesList');
            filesList.innerHTML = '';
            files.forEach(file => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td><input type="checkbox" value="${file.id}"></td>
                    <td>${file.filename}</td>
                    <td>${file.filesize}</td>
                    <td>${file.path}</td>
                    <td>${file.time}</td>
                `;
                filesList.appendChild(row);
            });

            currentPage = page;
            totalPages = response.totalPages; // 假设后端返回总页数
            updatePaginationControls();
        } else {
            alert("获取文件失败");
        }
    };
    xhr.send();
}

function updatePaginationControls() {
    document.getElementById("pagination").innerHTML = currentPage;
    document.getElementById("totalPage").innerHTML = totalPages;
    document.getElementById("prevPage").disabled = currentPage <= 1;
    document.getElementById("nextPage").disabled = currentPage >= totalPages;
    if(currentPage <= 1){
        document.getElementById("prevPage").style.display = "none";
    }else {
        document.getElementById("prevPage").style.display = "";
    }
    if(currentPage >= totalPages){
        document.getElementById("nextPage").style.display = "none";
    }else {
        document.getElementById("nextPage").style.display = "";
    }
}

function goToPrevPage() {
    if (currentPage > 1) {
        fetchFiles(currentPage - 1);
    }
}

function goToNextPage() {
    if (currentPage < totalPages) {
        fetchFiles(currentPage + 1);
    }
}

document.addEventListener("DOMContentLoaded", function() {
    fetchFiles(1); // 默认加载第一页
});

function deleteFiles() {
    // 收集所有选中的复选框
    let selectedFiles = document.querySelectorAll('input[type="checkbox"]:checked');
    if(selectedFiles.length == 0){
        alert("请选择需要删除的文件！")
    }else{
        let fileIds = [];
        selectedFiles.forEach(function(checkbox) {
            fileIds.push(checkbox.value);
        });

        console.log("删除文件的ID: ", fileIds.join(", "));
        // 实际操作：发送删除请求到后端，然后移除这些行或刷新列表
        // 创建XMLHttpRequest对象
        var xhr = new XMLHttpRequest();
        // 配置POST请求
        xhr.open("get", "/file/deleteFiles?fileIds="+fileIds, true);
        // 设置请求头以发送JSON格式的数据
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        // 定义请求完成的处理函数
        xhr.onload = function () {
            if (xhr.status >= 200 && xhr.status < 300) {
                // 请求成功，可以在这里更新UI或通知用户
                alert("删除成功！")
                // 可以在这里调用fetchFiles()来刷新文件列表
                fetchFiles(currentPage);
            } else {
                // 处理请求失败的情况
                console.error("Failed to delete files");
                alert("删除请求失败！")
            }
        };
        // 发送请求，将文件ID数组转换为JSON字符串
        xhr.send();
    }

}

function downloadFiles() {
    let selectedFiles = document.querySelectorAll('input[type="checkbox"]:checked');
    if(selectedFiles.length == 0){
        alert("请选择需要下载的文件！")
    }else {
        let fileIds = [];
        selectedFiles.forEach(function(checkbox) {
            fileIds.push(checkbox.value);
        });

        console.log("下载文件的ID： ", fileIds.join(", "));
        // 实际操作：发送下载请求到后端
        var downloadUrl = `/file/downloadFile?fileIds=` + fileIds;
        window.location.href = downloadUrl;
    }
}

function uploadFiles(){
    // 创建XMLHttpRequest对象
    var xhr = new XMLHttpRequest();
    var fileInput = document.getElementById('ulFile');
    var file = fileInput.files[0]; // 获取选择的文件

    // 创建FormData对象
    var formData = new FormData();
    formData.append('file', file); // 将文件添加到FormData对象中
    // 配置POST请求
    xhr.open("post", "/file/upload", true);
    // 设置请求头以发送JSON格式的数据
    // xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    // 定义请求完成的处理函数
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            // 请求成功，可以在这里更新UI或通知用户
            alert("上传成功！")
            // 可以在这里调用fetchFiles()来刷新文件列表
            fetchFiles(currentPage);
        } else {
            // 处理请求失败的情况
            console.error("Failed to delete files");
            var response = JSON.parse(xhr.responseText);
            if(response.msg != null){
                alert(response.msg);
            }else {
                alert("上传失败！")
            }
        }
    };
    xhr.send(formData);
}


