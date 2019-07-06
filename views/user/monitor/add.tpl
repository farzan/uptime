
<a class="btn btn-secondary mb-5" href="/user">Back</a>

<div class="row">
    <div class="col">
        <form method="post" action="/user/monitor/add">
            <div class="form-group">
                <label>Title</label>
                <input type="text" name="title" value="{{.monitor.Title}}" class="form-control">
            </div>

            <div class="form-group">
                <label>Method</label>
                <input type="text" name="method" class="form-control" value="GET" readonly="readonly">
            </div>

            <div class="form-group">
                <label>URL</label>
                <input type="text" name="url" value="{{.monitor.Url}}" class="form-control">
            </div>

            <div class="form-group">
                <label>Interval</label>
                <div class="custom-control custom-radio">
                    <input type="radio" id="interval1" name="interval" value="5" class="custom-control-input">
                    <label class="custom-control-label" for="interval1">5 seconds</label>
                </div>

                <div class="custom-control custom-radio">
                    <input type="radio" id="interval2" name="interval" value="5" class="custom-control-input">
                    <label class="custom-control-label" for="interval2">10 seconds</label>
                </div>

                <div class="custom-control custom-radio">
                    <input type="radio" id="interval3" name="interval" value="5" class="custom-control-input">
                    <label class="custom-control-label" for="interval3">30 seconds</label>
                </div>

                <div class="custom-control custom-radio">
                    <input type="radio" id="interval4" name="interval" value="5" class="custom-control-input">
                    <label class="custom-control-label" for="interval4">60 seconds</label>
                </div>
            </div>

            <div class="form-row">
                <div class="col-md-4 mb-3">
                    <label>Ok threshold</label>
                    <input type="text" name="ok" value="{{.monitor.Thresholds.GetOk.Seconds}}" class="form-control">
                    <div class="valid-tooltip">
                        Looks good!
                    </div>
                </div>

                <div class="col-md-4 mb-3">
                    <label>Warning threshold</label>
                    <input type="text" name="warning" value="{{.monitor.Thresholds.GetWarning.Seconds}}" class="form-control">
                    <div class="valid-tooltip">
                        Looks good!
                    </div>
                </div>

                <div class="col-md-4 mb-3">
                    <label>Critical threshold</label>
                    <input type="text" name="critical" value="{{.monitor.Thresholds.GetCritical.Seconds}}" class="form-control">
                    <div class="valid-tooltip">
                        Looks good!
                    </div>
                </div>
            </div>

            <div class="form-group">
                <input type="submit" value="Add" class="btn btn-primary">
            </div>
        </form>
    </div>
</div>