<!DOCTYPE html>

<html>
<head>
  <title>爬虫任务系统</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>

<body>

  <header>
    <div class="hed"><img class="hed" src="/static/img/logo.png"></div>
    <div class="had">爬虫任务系统</div>
  </header>
  <section>
    <ul>
      <li id="sy1" class="action" onclick="change(1)">首页</li>
      <li id="now1" onclick="change(2)">任务列表</li>
      <li id="new1" onclick="change(3)">添加任务</li>
    </ul>
    <!--iframe src="http://www.baidu.com"></iframe-->
    <!--首页-->
    <div id="sy" class="sec" style="display:block;">
      <div class="t2">欢迎使用爬虫任务系统</div>
      <div class="t3">使&nbsp用&nbsp说&nbsp明</div>
      <div class="t4">
        
      </div>
    </div>
    <!-----------进行中的-------------------------------------->
    <div id="now" class="sec" style="display:none;">
      <div class="t2">任务列表</div>
      <div class="table">
        <ul>
          <li class="action">任务名称</li>
          {{range .Task}}
            <li>{{.Name}}</li>
          {{end}}
        </ul>
        <ul>
          <li class="action">任务状态</li>
          {{range .Task}}
            {{if eq .Status 1}}
              <li class="on">进行中</li>
            {{else}}
              <li class="off">停止</li>
            {{end}}

          {{end}}
        </ul>
        <ul>
          <li class="action">任务操作</li>
          {{range .Task}}
            {{if eq .Status 1}}
              <li><button class="off" onclick="stopt({{.Name}})">停止任务</button></li>
            {{else}}
              <li><button class="on" onclick="startt({{.Name}})">开启任务</button></li>
            {{end}}
          {{end}}
        </ul>
        <ul>
          <li class="action">任务详情</li>
          {{range .Task}}
            <li><a onclick="digJump1({{.Name}},{{.Script}},{{.Time}},{{.Command}},{{.Des}})">点击详情</a></li>
          {{end}}
        </ul>
      </div>
    </div>
    <!---------添加任务---------------------------------------------->
    <div id="new" class="sec" style="display:none;">
      
      <div class="from">
        <div class="t2">添加任务</div>
        <p><label>任务名称：</label></label><input id="n1" type="text" name="name"/></p>
        <p><label>执行时间：</label><input id="t1" type="text" name="time"/></p>
        <p><label>返回类型：</label>
          <select id="s" name="script">
              <option value="html">html</option>
              <option value="json">json</option>
              <option value="jsonp">jsonp</option>
              <option value="xml">xml</option>
              <option value="text">text</option>
          </select>
        </p>
        <p><label>执行线程数：</label><input id="o1" type="text" name="order"/></p>
        <p><label>任务描述：</label><textarea id="d1"></textarea></p>

        <hr/>
        <div class="t2">添加规则</div>
        <p><label>项目名称：</label></label><input id="in1" type="text" name="in1"/></p>
        <p><label>项目规则：</label><input id="ir1" type="text" name="ir1"/></p>
        <button id="1" class="add-rule"  onclick="addRule()">+</button>
        <hr/>
        <div><input class="submit" type="button" name="submit" value="提交" onclick="submitTask(1)" /></div>
      </div>

    </div>
    
  </section>
  <footer>
    
  </footer>
  <!-------------对话框-------------------------->
  <dialog id="dig1" open class="dig">
    <button onclick="digClose()">X</button>
    <div id="t1" class="t2"></div>
    <div class="from">
        <p>任务名称：<input id="n" type="text" name="name"/></p>
        <p>执行时间：<input id="t" type="text" name="time"/></p>
        <p>返回类型：
          <select id="s" name="script">
              <option value="html">html</option>
              <option value="json">json</option>
              <option value="jsonp">jsonp</option>
              <option value="xml">xml</option>
              <option value="text">text</option>
          </select>
        </p>
        <p>脚本名称：<input id="o" type="text" name="order"/><font color="#ff0000">请填写含路径脚本名称</font></p>
        <p><span>任务描述：</span><textarea id="d"></textarea></p>
        <div><input type="button" name="submit" value="提交" onclick="submitTask(2)" /></div>
    </div>
  </dialog>

  <script type="text/javascript" src="/static/js/jquery-1.9.0.min.js"></script>
  <script type="text/javascript">

    function change(flag){
      if(flag==1){
        $("#sy").show();
        $("#now").hide();
        $("#new").hide();
        $("#conf").hide();
        $("#sy1").attr("class","action");
        $("#now1").attr("class","");
        $("#new1").attr("class","");
        $("#conf1").attr("class","");
      }else if(flag==2){
        $("#sy").hide();
        $("#now").show();
        $("#new").hide();
        $("#conf").hide();
        $("#sy1").attr("class","");
        $("#now1").attr("class","action");
        $("#new1").attr("class","");
        $("#conf1").attr("class","");
      }else if(flag==3){
        $("#sy").hide();
        $("#now").hide();
        $("#new").show();
        $("#conf").hide();
        $("#sy1").attr("class","");
        $("#now1").attr("class","");
        $("#new1").attr("class","action");
        $("#conf1").attr("class","");
      }else{
        $("#sy").hide();
        $("#now").hide();
        $("#new").hide();
        $("#conf").show();
        $("#sy1").attr("class","");
        $("#now1").attr("class","");
        $("#new1").attr("class","");
        $("#conf1").attr("class","action");
      }
    }
    function addRule(){
      var num=$(".add-rule").attr("id");
      console.log(num);
      num=(parseInt(num)+1).toString();
      console.log(num);
      $(".add-rule").before("<hr class=\"in\"/><p><label>项目名称：</label><input id=\"in"+num+"\" type=\"text\" name=\"in"+num+"\"/><p><p><label>项目规则：</label><input id=\"ir"+num+"\" type=\"text\" name=\"ir"+num+"\"/></p>");
      $(".add-rule").attr("id",num);
    }
    function digJump1(n,s,t,o,d){
      $("#dig2").hide();
      $("#dig1").show();
      $("#t1").text(n+"任务详情信息");
      $("#n").attr("value",n);
      $("#o").attr("value",o);
      $("#t").attr("value",t);
      $("#d").text(d);
      var ops=$("#s")[0].options;
      for (k in ops){
        if(ops[k].label==s)
          ops[k].selected=true;
      }
    }
    function digJump2(f,n,o){
      $("#dig1").hide();
      $("#dig2").show();
      if(f==1){
        $("#t2").text("添加配置");
      }else{
        $("#t2").text("修改"+n+"脚本信息");
      }
      $("#n11").attr("value",n);
      $("#n2").attr("value",o);
    }
    function digClose(){
      $("#dig1").hide();
      $("#dig2").hide();
    }
    function submitTask(f){
      if(f==1){
        var n=$("#n1")[0].value;
        var o=$("#o1")[0].value;
        var t=$("#t1")[0].value;
        var d=$("#d1")[0].value;
        var s=$("#s1").find("option:selected").text();
        alert(n+o+t+d+s);
        $.get('/addTask',{"n":n,"o":o,"t":t,"s":s,"d":d},function(data){
          if(data.code==1){
            alert("操作成功！！");
          }else{
            alert("操作失败！！");
          }
          location.reload() 
        });
      }else{
        var n=$("#n")[0].value;
        var o=$("#o")[0].value;
        var t=$("#t")[0].value;
        var d=$("#d")[0].value;
        var s=$("#s").find("option:selected").text();
        $.get('/updateTask',{"n":n,"o":o,"t":t,"s":s,"d":d},function(data){
          $("#dig1").hide();
          if(data.code==1){
            alert("操作成功！！");
          }else{
            alert("操作失败！！");
          }
          location.reload() 
        });
      }
    }
    function submitConf(){
      var n=$("#n11")[0].value;
      var o=$("#n2")[0].value;
      $.get('/updateConf',{"n":n,"o":o},function(data){
        $("#dig2").hide();
        if(data.code==1){
          alert("操作成功！！");
        }else{
          alert("操作失败！！");
        }
        location.reload() 
      });
    }
    function stopt(n){
      $.get('/stop',{"name":n},function(data){
        if(data.code==1){
          alert("操作成功！！");
        }else{
          alert("操作失败！！");
        }
        location.reload() 
      });
    }
    function startt(n){
      $.get('/start',{"name":n},function(data){
        if(data.code==1){
          alert("操作成功！！");
        }else{
          alert("操作失败！！");
        }
        location.reload() 
      });
    }
  </script>
  <!--script type="text/javascript" src="http://cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script-->
</body>
</html>
