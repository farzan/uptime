<!DOCTYPE html>

<html>
<head>
  <title>Uptime Monitor</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

</head>

<body>
    <div>
        <form method="POST" action="/user/login">
            <div class="form-group">
                <label>
                    Username:
                    <input class="form-control" type="text" name="username">
                </label>
            </div>

            <div class="form-group">
                <label>
                    Password:
                    <input class="form-control" type="text" name="password">
                </label>
            </div>

            <input class="btn btn-primary", type="submit" value="Login">

            <br>
            <br>
            <br>
            <a href="/user">User monitors</a>
        </form>
    </div>
</body>
</html>
