<div class="row">
    <div class="col">
        <form method="post" action="/user/monitor/{{.monitor.Id}}/delete">
            <div class="alert alert-danger">
                Are you sure?
            </div>

            <a class="btn btn-secondary mb-5" href="/user">Back</a>
            <button type="submit" class="btn btn-primary mb-5" href="/user">Delete</button>
        </form>
    </div>
</div>