<a class="btn btn-secondary mb-5" href="/user">Back</a>

<h3 class=" mb-3">Count: {{.count}}</h3>

<div class="row">
    <div class="col">
        <table class="table">
            <thead>
                <tr>
                    <th scope="col">Date</th>
                    <th scope="col">Duration</th>
                    <th scope="col">Status code</th>
                    <th scope="col">Description</th>
                </tr>
            </thead>
            <tbody>
            {{range $key, $response := .responses}}
                {{if not $response.IsError}}
                    <tr class="{{if $response.IsOk}}table-success{{else if $response.IsWarning}}table-warning{{else}}table-danger{{end}}">
                        <td>{{$response.Start.Year}}-{{$response.Start.Month}}-{{$response.Start.Day}} {{$response.Start.Hour}}:{{$response.Start.Minute}}:{{$response.Start.Second}}</td>
                        <td>{{printf "%.3f" $response.Duration.Seconds}}s</td>
                        <td>{{$response.StatusCode}}</td>
                        <td>{{if $response.IsOk}}Ok{{else if $response.IsWarning}}Warning{{else}}Critical!{{end}}</td>
                    </tr>
                {{else}}
                    <tr class="bg-danger text-white">
                        <td>{{$response.Start.Year}}-{{$response.Start.Month}}-{{$response.Start.Day}} {{$response.Start.Hour}}:{{$response.Start.Minute}}:{{$response.Start.Second}}</td>
                        <td>{{printf "%.3f" $response.Duration.Seconds}}s</td>
                        <td>{{$response.StatusCode}}</td>
                        <td>{{$response.Error}}</td>
                    </tr>
                {{end}}
            {{end}}
            </tbody>
        </table>
    </div>
</div>