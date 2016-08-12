var sid

    function show(){
      $("#postcanshu").show()
    }
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
    function addRule(f){
      var num=$(".add-rule").attr("id");
      num=(parseInt(num)+1).toString();
      if(f==1){
        $(".add-rule").before("<hr class=\"in\"/><p><label>项目名称：</label><input id=\"in"+num+"\" type=\"text\" name=\"in"+num+"\" placeholder=\"抓取字段名称\"/><p><p><label>项目规则：</label><input id=\"ir"+num+"\" type=\"text\" name=\"ir"+num+"\"/ placeholder=\"选择器规则 比如 .in p\"></p><p><label>项目函数：</label><input id=\"if"+num+"\" type=\"text\" name=\"if"+num+"\" placeholder=\"text或者属性名称 比如：href\"/></p>");
      }else{
        $(".add-rule").before("<hr class=\"in\"/><p><label>项目名称：</label><input id=\"in1"+num+"\" type=\"text\" name=\"in1"+num+"\" placeholder=\"抓取字段名称\"/><p><p><label>项目规则：</label><input id=\"ir1"+num+"\" type=\"text\" name=\"ir1"+num+"\"/ placeholder=\"选择器规则 比如 .in p\"></p><p><label>项目函数：</label><input id=\"if1"+num+"\" type=\"text\" name=\"if1"+num+"\" placeholder=\"text或者属性名称 比如：href\"/></p>");
      }
      $(".add-rule").attr("id",num);
    }
    function digJump1(id,taskname, theardnum,cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun){
      $("#dig").css("display","block");
      sid=id
      $("#taskname1").attr("value",taskname);
      $("#theardnum1").attr("value",theardnum);
      $("#cron1").attr("value",cron);
      $("#des1").text(des);
      var ops=$("#dbtype1")[0].options;
      for (k in ops){
        if(ops[k].label==dbtype)
          ops[k].selected=true;
      }
      $("#dbhost1").attr("value",dbhost);
      $("#dbportq").attr("value",dbport);
      $("#dbname1").attr("value",dbname);
      $("#dbuser1").attr("value",dbuser);
      $("#dbpasswd1").attr("value",dbpasswd);
      var ops=$("#texttype1")[0].options;
      for (k in ops){
        if(ops[k].label==texttype)
          ops[k].selected=true;
      }
      var ops=$("#reqtype1")[0].options;
      for (k in ops){
        if(ops[k].label==reqtype)
          ops[k].selected=true;
      }
      $("#rooturl1").attr("value",rooturl);
      $("#cookie1").attr("value",cookie);
      $("#headerfile1").attr("value",headerfile);
      var ops=$("#useproxy1")[0].options;
      for (k in ops){
        if(ops[k].label==useproxy)
          ops[k].selected=true;
      }
      $("#pagepre1").attr("value",pagepre);
      $("#pagerule1").attr("value",pagerule);
      $("#pagefun1").attr("value",pagefun);
      $("#postdata1").attr("value",postdata);
      
      

    }
    
    function digClose(){
      $("#dig").hide();
    }
    function submitTask(f){
      if(f==1){
        var taskname=$("#taskname")[0].value;
        var cron=$("#cron")[0].value;
        var theardnum=$("#theardnum")[0].value;
        var des=$("#des")[0].value;
        var dbtype=$("#dbtype").find("option:selected").text();
        var dbhost=$("#dbhost")[0].value;
        var dbport=$("#dbport")[0].value;
        var dbname=$("#dbname")[0].value;
        var dbuser=$("#dbuser")[0].value;
        var dbpasswd=$("#dbpasswd")[0].value;
        var texttype=$("#texttype").find("option:selected").text();
        var reqtype=$("#reqtype").find("option:selected").text();
        var postdata=$("#postdata")[0].value;
        var rooturl=$("#rooturl")[0].value;
        var cookie=$("#cookie")[0].value;
        var headerfile=$("#headerfile")[0].value;
        var useproxy=$("#useproxy").find("option:selected").text();
        var pagepre=$("#pagepre")[0].value;
        var pagerule=$("#pagerule")[0].value;
        var pagefun=$("#pagefun")[0].value;
        var ii=1;
        var rule = new Array();
        var fun=new Array();
        var name=new Array();
        while($("#in"+ii.toString()).length==1){
          if($("#in"+ii.toString())[0].value!=""){
            rule.push($("#ir"+ii.toString())[0].value);
            fun.push($("#if"+ii.toString())[0].value);
            name.push($("#in"+ii.toString())[0].value);
          }
          ii++;
        }
        name=name.join("|");
        rule=rule.join("|");
        fun=fun.join("|");

        $.get('/addconf',{"taskname":taskname,"cron":cron,"des":des,"dbtype":dbtype,"dbhost":dbhost,"dbport":dbport,"dbname":dbname,"dbuser":dbuser,"dbpasswd":dbpasswd,"reqtype":reqtype,"rooturl":rooturl,"cookie":cookie,"headerfile":headerfile,"useproxy":useproxy,"texttype":texttype,"postdata":postdata,"pagepre":pagepre,"pagerule":pagerule,"pagefun":pagefun,"theardnum":theardnum,"name":name,"rule":rule,"fun":fun},function(data){
          if(data.code==1){
            alert("操作成功！！");
          }else{
            alert("操作失败！！");
          }
          location.reload() 
        });
      }else{
        var taskname=$("#taskname1")[0].value;
        var cron=$("#cron1")[0].value;
        var theardnum=$("#theardnum1")[0].value;
        var des=$("#des1")[0].value;
        var dbtype=$("#dbtype1").find("option:selected").text();
        var dbhost=$("#dbhost1")[0].value;
        var dbport=$("#dbport1")[0].value;
        var dbname=$("#dbname1")[0].value;
        var dbuser=$("#dbuser1")[0].value;
        var dbpasswd=$("#dbpasswd1")[0].value;
        var texttype=$("#texttype1").find("option:selected").text();
        var reqtype=$("#reqtype1").find("option:selected").text();
        var postdata=$("#postdata1")[0].value;
        var rooturl=$("#rooturl1")[0].value;
        var cookie=$("#cookie1")[0].value;
        var headerfile=$("#headerfile1")[0].value;
        var useproxy=$("#useproxy1").find("option:selected").text();
        var pagepre=$("#pagepre1")[0].value;
        var pagerule=$("#pagerule1")[0].value;
        var pagefun=$("#pagefun1")[0].value;
        var ii=1;
        var rule = new Array();
        var fun=new Array();
        var name=new Array();
        while($("#in1"+ii.toString()).length==1){
          if($("#in1"+ii.toString())[0].value!=""){
            rule.push($("#ir1"+ii.toString())[0].value);
            fun.push($("#if1"+ii.toString())[0].value);
            name.push($("#in1"+ii.toString())[0].value);
          }
          ii++;
        }
        name=name.join("|");
        rule=rule.join("|");
        fun=fun.join("|");
        alert(taskname+cron+theardnum+des+dbtype);
        // $.get('/updateTask',{"n":n,"o":o,"t":t,"s":s,"d":d},function(data){
        //   $("#dig1").hide();
        //   if(data.code==1){
        //     alert("操作成功！！");
        //   }else{
        //     alert("操作失败！！");
        //   }
        //   location.reload() 
        // });
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
    function stopt(name){
      $.get('/stop',{"name":name},function(data){
        if(data.code==1){
          alert("操作成功！！");
        }else{
          alert("操作失败！！");
        }
        location.reload() 
      });
    }
    function startt(id,name,cron){
      $.get('/start',{"id":id,"name":name,"cron":cron},function(data){
        if(data.code==1){
          alert("操作成功！！");
        }else{
          alert("操作失败！！");
        }
        location.reload() 
      });
    }