
<div class="row">
    <div class="col">
        <div>
            <h2 class="float-left mb-5">Hello {{.user.Name}}</h2>
            <a class="btn btn-outline-dark btn-sm float-right mt-2" href="">Update your account</a>
        </div>
    </div>
</div>

<div class="row">
    <div class="col">
        <h3 class="mb-4">Your monitors:</h3>
        <ol class="list-group list-group-flush mb-3">
            {{range $key, $monitor := .monitors}}
            <li class="list-group-item list-group-item-action" href="#">
                <div class="float-left">
                    <h4>{{$monitor.Title}}</h4>
                    <div>
                        <strong>{{$monitor.Method}}</strong> <code>{{$monitor.Url}}</code>
                    </div>
                    <div>Interval: <code>{{$monitor.Interval}} seconds</code></div>
                    <div>
                        Thresholds:
                        Ok: <code>{{$monitor.Thresholds.Ok}}</code>
                        Warning: <code>{{$monitor.Thresholds.Warning}}</code>
                        Critical: <code>{{$monitor.Thresholds.Critical}}</code>
                    </div>
                </div>
                <div class="float-right">
                    <a href="/monitor/report/{{$monitor.Id}}">report</a>
                    |
                    <a href="/monitor/update/{{$monitor.Id}}">update</a>
                    |
                    <a href="/monitor/delete/{{$monitor.Id}}">delete</a>
                </div>
            </li>
            {{end}}
        </ol>
        <div>
            <a class="btn btn-dark" href="#">Add new monitor</a>
        </div>
    </div>
</div>
