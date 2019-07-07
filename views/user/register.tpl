<!DOCTYPE html>

<html>
<head>
  <title>Uptime Monitor</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

</head>

<body>
    <div class="row">
        <div class="col-6">
            <form method="POST" action="/user/register">
                <div class="form-group">
                    <label>
                        Email:
                        <input class="form-control" type="text" name="email">
                    </label>
                </div>

                <div class="form-group">
                    <label>
                        Password:
                        <input class="form-control" type="password" name="password">
                    </label>
                </div>

                <div class="form-group">
                    <label>
                        Password repeat:
                        <input class="form-control" type="password" name="password-repeat">
                    </label>
                </div>

                <button class="btn btn-primary" type="submit" value="Login">Register</button>
                or <a href="/user/login">Login</a>
            </form>
        </div>
    </div>
</body>
</html>
