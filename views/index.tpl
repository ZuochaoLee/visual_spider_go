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
        <h2><font color="#03a">|</font>系统简介</h2>
        <p>系统是基于golang开发的可视化爬虫系统，目标是以最小的代价最高的运行效率完成垂直领域的爬虫任务。</p>
        <h2><font color="#03a">|</font>系统架构</h2>
        <p>系统由爬虫部分，任务调度部分，数据交互部分组成。</p>
        <p>爬虫部分是系统的核心，由下载器，任务队列，页面解析器，数据存储通道四部分组成，其中页面解析器和数据存储通道可以由用户配置使用，也可以进行二次开发来满足具体任务需求，据大多数情况下，系统的模版均可覆盖，用户只需配置一些参数即可使用，无需编写代码。</p>
        <p>任务调度部分是用golang模仿cron实现的，支持cron语法，可以在系统内很容易实现定时任务调度</p>
        <h2><font color="#03a">|</font>操作说明</h2>
        <h3><font color="#0aa">•</font>添加任务</h3>
        <p>添加任务是配置抓取中用到的关键信息，根据对垂直领域爬虫的抽象，分为六个部分</p>
        <p>1、任务信息：包括 任务名称，执行时间的cron表达式，执行线程数，任务的描述等</p>
        <p>2、存储数据库信息：包括 数据库类型（设计支持mysql，mongodb，redis,file,console等）主机名，端口号，库名，用户名，密码等信息</p>
        <p>3、下载器配置信息：包括 返回类型，请求方法，入口URL，cookie设置，hearder设置，使用代理等</p>
        <p>4、翻页规则：支持最多四层翻页变量设置，最内层必须为页码（格式为：1 1 100 即 起始页 步长 终止页），外层为自定义列表（列表项用空格隔开即可）</p>
        <p>5、种子解析规则：从信息列表页获取详情页链接种子的相关规则，包括css选择器唯一规则，选取属性类型，以及如果是相对路径需要的前缀URL等</p>
        <p>6、页面规则：从详情页解析相关信息规则，包括词项名称，css选择器唯一规则，选取属性类型，可添加多个</p>
        <h3><font color="#0aa">•</font>任务启停</h3>
        <p>在任务列表中可以操纵启停任务，任务启动不是及时运行，是把任务添加到了定时任务系统中，同样，停止任务只是从定时任务系统中删除了任务，不是及时停止任务。由于系统效率超群，大多数任务都会在十分钟以内，及时停止不是那么必须</p>
        <h3><font color="#0aa">•</font>任务修改</h3>
        <p>如果遇到被爬网站页面规则改版，可以在任务列表点击任务详情修改任务配置，具体操作同任务添加</p>
        <h3><font color="#0aa">•</font>配置信息</h3>
        <p>dbuser ：项目需要数据库用户名</p>
        <p>dbname ：项目需要数据库名</p>
        <p>dbpassword ：项目需要数据库密码</p>
        <p>dbhost ：项目需要数据库主机IP</p>
        <p>dbport ：项目需要数据库端口号</p>
        <p>iphost ：IP代理redis主机IP</p>
        <p>ipport ：IP代理redis端口</p>
        <p>ippassword ：IP代理redis密码</p>
        <p>ipdb ：IP代理redis数据库号</p>
        <p>其他配置参考beego配置参数</p>
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
