<!DOCTYPE html>

<html>
<head>
  <title>Uptime Monitor</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

</head>

<body>
    <div class="row">
        <div class="col-6">
            <form method="POST" action="/user/login">
                <div class="form-group">
                    <label>
                        Username:
                        <input class="form-control" type="text" name="email">
                    </label>
                </div>

                <div class="form-group">
                    <label>
                        Password:
                        <input class="form-control" type="password" name="password">
                    </label>
                </div>

                <button class="btn btn-primary" type="submit" value="Login">Login</button>
                or <a href="/user/register">Register</a>

                <br>
                <br>
                <br>
                <a href="/user">User monitors</a>
            </form>
        </div>
    </div>
</body>
</html>
