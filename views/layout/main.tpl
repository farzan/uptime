<html>
    <head>
        <script src="/static/bootstrap/js/bootstrap.min.js"></script>
        <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css" >
    </head>

    <body>
        <nav class="navbar navbar-expand-lg navbar-light bg-light mb-4">
            <a class="navbar-brand" href="#">Uptime Monitor</a>
            {{if .isLoggedIn}}
            <form class="form-inline my-2 my-lg-0" method="post" action="/user/logout">
                <button class="btn btn-dark my-2 my-sm-0" type="submit">Logout</button>
            </form>
            {{end}}
        </nav>

        <div class="container">
            <div class="row">
                <div class="col">
                    <div class="alert alert-danger">
                        {{.flash.error}}
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col mb-5">
                    {{.LayoutContent}}
                </div>
            </div>
        </div>
    </body>
</html>