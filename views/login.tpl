<DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    <title>登录</title>
    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
  </head>

  <body>

    <div class="container" style="padding-top: 10%">
      <div class="col-md-6 col-md-offset-3">

      <form class="form-signin" action="" method="post">
        <h2 class="form-signin-heading">Please sign in</h2>
        <input type="text" id="inputEmail" name="username" class="form-control" placeholder="Username" required autofocus>
        <input type="password" id="inputPassword" name="password" class="form-control" placeholder="Password" required>
        <div class="checkbox">

        </div>
        <input class="btn btn-lg btn-primary btn-block" type="submit" value="Sign in">
      </form>
    </div>
    <p>{{ .flash.error }}</p>
    </div> <!-- /container -->
    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
  </body>
</html>


