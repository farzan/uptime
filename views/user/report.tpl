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
                <tr class="{{if $response.IsOk}}table-success{{else if $response.IsWarning}}table-warning{{else}}table-danger{{end}}">
                    <td>{{$response.Start.Year}}-{{$response.Start.Month}}-{{$response.Start.Day}} {{$response.Start.Hour}}:{{$response.Start.Minute}}:{{$response.Start.Second}}</td>
                    <td>{{$response.DurationInSeconds}}</td>
                    <td>{{$response.StatusCode}}</td>
                    <td>ok?</td>
                </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>