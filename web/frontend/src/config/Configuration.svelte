<script>
    import { Button, Card, CardTitle } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import { fade } from "svelte/transition";
</script>

<Card>
    <CardTitle class="mb-3">InfluxDB Configuration</CardTitle>
    <div class="mb-3">
        <label for="database" class="form-label">Database (Bucket)</label>
        <input type="text" class="form-control" id="database" name="database" />

        <label for="host" class="form-label">Host</label>
        <input type="text" class="form-control" id="host" name="host" />

        <label for="port" class="form-label">Port</label>
        <input type="text" class="form-control" id="port" name="port" />

        <label for="user" class="form-label">User</label>
        <input type="text" class="form-control" id="user" name="user" />

        <label for="password" class="form-label">Password</label>
        <input
            type="password"
            class="form-control"
            id="password"
            name="password"
        />

        <label for="organization" class="form-label">Organization</label>
        <input
            type="text"
            class="form-control"
            id="organization"
            name="organization"
        />

        <label for="ssl" class="form-label">SSL</label>
        <input type="checkbox" class="form-control" id="ssl" name="ssl" />

        <label for="batch_size" class="form-label">Batch Size</label>
        <input
            type="number"
            class="form-control"
            id="batch_size"
            name="batch_size"
        />

        <label for="retry_interval" class="form-label">Retry Interval</label>
        <input
            type="text"
            class="form-control"
            id="retry_interval"
            name="retry_interval"
        />

        <label for="retry_exponential_base" class="form-label"
            >Retry Exponential Base</label
        >
        <input
            type="number"
            class="form-control"
            id="retry_exponential_base"
            name="retry_exponential_base"
        />

        <label for="max_retries" class="form-label">Max Retries</label>
        <input
            type="number"
            class="form-control"
            id="max_retries"
            name="max_retries"
        />

        <label for="max_retry_time" class="form-label">Max Retry Time</label>
        <input
            type="text"
            class="form-control"
            id="max_retry_time"
            name="max_retry_time"
        />

        <label for="meta_as_tags" class="form-label">Meta as Tags</label>
        <input
            type="text"
            class="form-control"
            id="meta_as_tags"
            name="meta_as_tags"
        />
    </div>
</Card>
