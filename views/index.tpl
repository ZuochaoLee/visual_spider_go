<!DOCTYPE html>

<html>
<head>
  <title>可视化爬虫控制台</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>

<body>

  <header>
    <div class="hed"><img class="hed" src="/static/img/logo.png"></div>
    <div class="had">可视化爬虫控制台</div>
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
          {{range .Conf}}
            <li>{{.TaskName}}</li>
          {{end}}
        </ul>
        <ul>
          <li class="action">任务状态</li>
          {{range .Conf}}
            {{if eq .Status 1}}
              <li class="on">进行中</li>
            {{else}}
              <li class="off">停止</li>
            {{end}}

          {{end}}
        </ul>
        <ul>
          <li class="action">任务操作</li>
          {{range .Conf}}
            {{if eq .Status 1}}
              <li><button class="off" onclick="stopt({{.TaskName}})">停止任务</button></li>
            {{else}}
              <li><button class="on" onclick="startt({{.ID}},{{.TaskName}},{{.Cron}})">开启任务</button></li>
            {{end}}
          {{end}}
        </ul>
        <ul>
          <li class="action">任务详情</li>
          {{range .Conf}}
          
            <li><a onclick="digJump1({{.ID}},{{.TaskName}},{{.Theardnum}},{{.Cron}},{{.Des}},{{.Dbtype}},{{.Dbhost}},{{.Dbport}},{{.Dbname}},{{.Dbuser}},{{.Dbpasswd}},{{.ReqType}},{{.RootUrl}},{{.Cookie}},{{.HeaderFile}},{{.UseProxy}},{{.TextType}},{{.PostData}},{{.PagePre}},{{.PageRule}},{{.PageFun}},{{.PageFour}},{{.PageThree}},{{.PageTwo}},{{.PageOne}})">点击详情</a></li>
          {{end}}
        </ul>
      </div>
    </div>
    <!---------添加任务---------------------------------------------->
    <div id="new" class="sec" style="display:none;">
      
      <div class="from">
        <div class="t2">添加任务</div>
        <p><label>任务名称：</label></label><input id="taskname" type="text" name="taskname"/></p>
        <p><label>执行时间：</label><input id="cron" type="text" name="cron" placeholder="cron表达式 "/></p>
        
        <p><label>执行线程数：</label><input id="theardnum" type="text" name="theardnum"/></p>
        <p><label>任务描述：</label><textarea id="des" name="des"></textarea></p>

        <div class="t2">存储数据库配置</div>
        <p><label>数据库类型：</label>
          <select id="dbtype" name="dbtype">
              <option value="mysql">mysql</option>
              <option value="mongodb">mongodb</option>
              <option value="redis">redis</option>
              <option value="file">file</option>
              <option value="console">console</option>
          </select>
        </p>
        <p><label>数据库主机：</label></label><input id="dbhost" type="text" name="dbhost"/></p>
        <p><label>数据库端口：</label><input id="dbport" type="text" name="dbport"/></p> 
        <p><label>数据库名称：</label><input id="dbname" type="text" name="dbname"/></p>
        <p><label>数据库用户：</label><input id="dbuser" type="text" name="dbuser"/></p>
        <p><label>数据库密码：</label><input id="dbpasswd" type="text" name="dbpasswd"/></p>

        <div class="t2">下载器配置</div>

        <p><label>返回类型：</label>
          <select id="texttype" name="texttype">
              <option value="html">html</option>
              <option value="json">json</option>
              <option value="jsonp">jsonp</option>
              <option value="xml">xml</option>
              <option value="text">text</option>
          </select>
        </p>
        <p><label>请求方法：</label>
          <select id="reqtype" name="reqtype" onchange="show()">
              <option value="GET">GET</option>
              <option value="POST">POST</option>
              
          </select>
        </p>
        <p id='postcanshu' style="display:none"><label>POST参数：</label><input id="postdata" type="text" name="postdata"/></p>
        <p><label>根URL：</label></label><input id="rooturl" type="text" name="rooturl"/></p>
        <p><label>cookie：</label><input id="cookie" type="text" name="cookie"/></p>
        
        <p><label>header文件路径：</label><input id="headerfile" type="text" name="headerfile"/></p>
        
        <p><label>是否使用代理：</label>
          <select id="useproxy" name="useproxy">
              <option value='0'>否</option>
              <option value="1">是</option>
              
          </select>
        </p>

        <div class="t2">翻页规则</div>

        <p><label>内一层：</label></label><input id="pagefour" type="text" name="pagefour" placeholder="只能是页码范围 格式如：起始 步长 终止 例：1 1 100"/></p>
        <p><label>内二层</label><input id="pagethree" type="text" name="pagethree" placeholder="列表空格隔开"/></p>
        <p><label>内三层</label><input id="pagetwo" type="text" name="pagetwo" placeholder="列表空格隔开"/></p>
        <p><label>内四层</label><input id="pageone" type="text" name="pageone" placeholder="列表空格隔开"/></p>

        <div class="t2">种子解析规则</div>

        <p><label>URL前缀：</label></label><input id="pagepre" type="text" name="pagepre" placeholder=""/></p>
        <p><label>种子规则：</label><input id="pagerule" type="text" name="pagerule" placeholder="选择器规则 比如 .in p"/></p>
        <p><label>种子函数：</label><input id="pagefun" type="text" name="pagefun" placeholder="text或者属性名称 比如：href"/></p>

        <div class="t2">添加页面规则</div>

        <p><label>项目名称：</label></label><input id="in1" type="text" name="in1" placeholder="抓取字段名称"/></p>
        <p><label>项目规则：</label><input id="ir1" type="text" name="ir1" placeholder="选择器规则 比如 .in p"/></p>
        <p><label>项目函数：</label><input id="if1" type="text" name="if1" placeholder="text或者属性名称 比如：href"/></p>
        <button id="1" class="add-rule d1"  onclick="addRule(1)">+</button>
        <hr/>
        <div><input class="submit" type="button" name="submit" value="提交" onclick="submitTask(1)" /></div>
      </div>

    </div>
    
  </section>
  <footer>
    
  </footer>
  <!-------------对话框-------------------------->
  <dialog id="dig"  class="dig">
    <button onclick="digClose()">X</button>
    <div id="t1" class="t2"></div>
    <div class="from">
        <div class="t2">修改任务</div>
        <p><label>任务名称：</label></label><input id="taskname1" type="text" name="taskname"/></p>
        <p><label>执行时间：</label><input id="cron1" type="text" name="cron" placeholder="cron表达式 "/></p>
        
        <p><label>执行线程数：</label><input id="theardnum1" type="text" name="theardnum"/></p>
        <p><label>任务描述：</label><textarea id="des1" name="des"></textarea></p>

        <div class="t2">存储数据库配置</div>
        <p><label>数据库类型：</label>
          <select id="dbtype1" name="dbtype">
              <option value="mysql">mysql</option>
              <option value="mongodb">mongodb</option>
              <option value="redis">redis</option>
              <option value="file">file</option>
              <option value="console">console</option>
          </select>
        </p>
        <p><label>数据库主机：</label></label><input id="dbhost1" type="text" name="dbhost"/></p>
        <p><label>数据库端口：</label><input id="dbport1" type="text" name="dbport"/></p> 
        <p><label>数据库名称：</label><input id="dbname1" type="text" name="dbname"/></p>
        <p><label>数据库用户：</label><input id="dbuser1" type="text" name="dbuser"/></p>
        <p><label>数据库密码：</label><input id="dbpasswd1" type="text" name="dbpasswd"/></p>

        <div class="t2">下载器配置</div>

        <p><label>返回类型：</label>
          <select id="texttype1" name="texttype">
              <option value="html">html</option>
              <option value="json">json</option>
              <option value="jsonp">jsonp</option>
              <option value="xml">xml</option>
              <option value="text">text</option>
          </select>
        </p>
        <p><label>请求方法：</label>
          <select id="reqtype1" name="reqtype" onchange="show()">
              <option value="GET">GET</option>
              <option value="POST">POST</option>
              
          </select>
        </p>
        <p id='postcanshu1' style="display:none"><label>POST参数：</label><input id="postdata1" type="text" name="postdata"/></p>
        <p><label>根URL：</label></label><input id="rooturl1" type="text" name="rooturl"/></p>
        <p><label>cookie：</label><input id="cookie1" type="text" name="cookie"/></p>
        
        <p><label>header文件路径：</label><input id="headerfile1" type="text" name="headerfile"/></p>
        
        <p><label>是否使用代理：</label>
          <select id="useproxy1" name="useproxy">
              <option value='0'>否</option>
              <option value="1">是</option>
              
          </select>
        </p>

        <div class="t2">翻页规则</div>

        <p><label>内一层：</label></label><input id="pagefour1" type="text" name="pagefour1" placeholder="只能是页码范围 格式如：起始 步长 终止 例：1 1 100"/></p>
        <p><label>内二层</label><input id="pagethree1" type="text" name="pagethree1" placeholder="列表空格隔开"/></p>
        <p><label>内三层</label><input id="pagetwo1" type="text" name="pagetwo1" placeholder="列表空格隔开"/></p>
        <p><label>内四层</label><input id="pageone1" type="text" name="pageone1" placeholder="列表空格隔开"/></p>

        <div class="t2">种子解析规则</div>

        <p><label>URL前缀：</label></label><input id="pagepre1" type="text" name="pagepre" placeholder=""/></p>
        <p><label>种子规则：</label><input id="pagerule1" type="text" name="pagerule" placeholder="选择器规则 比如 .in p"/></p>
        <p><label>种子函数：</label><input id="pagefun1" type="text" name="pagefun" placeholder="text或者属性名称 比如：href"/></p>

        <div class="t2">添加页面规则</div>

        <p><label>项目名称：</label></label><input id="in11" type="text" name="in1" placeholder="抓取字段名称"/></p>
        <p><label>项目规则：</label><input id="ir11" type="text" name="ir1" placeholder="选择器规则 比如 .in p"/></p>
        <p><label>项目函数：</label><input id="if11" type="text" name="if1" placeholder="text或者属性名称 比如：href"/></p>
        <button id="1" class="add-rule d2"  onclick="addRule(2)">+</button>
        <hr/>
        <div><input class="submit" type="button" name="submit" value="提交" onclick="submitTask(2)" /></div>
      </div>
    </div>
  </dialog>

  <script type="text/javascript" src="/static/js/jquery-1.9.0.min.js"></script>
  <script type="text/javascript" src="/static/js/main.js">

  </script>
  <!--script type="text/javascript" src="http://cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script-->
</body>
</html>
