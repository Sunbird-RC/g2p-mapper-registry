<script>
    import {getTransactionId} from "../util.js";
    import {config} from "../config.js";

    let logs = []
    setInterval(async () => {
        const res = await fetch(`${config.BASE_URL}/callback-service/v0.1.0/logs?transaction_id=${getTransactionId()}`);
        logs = (await res.json()).logs.reverse();
    }, 2000)

</script>
{#each logs as log}
    <div>
        <h4 style="margin-left: 1rem ">{JSON.parse(log).header.action} - {new Date(JSON.parse(log).header.message_ts).toLocaleString()} - Status {JSON.parse(log).header.message_ts}</h4>
        <pre>
            {log}
        </pre>
    </div>
{/each}
