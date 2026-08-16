<script lang="ts">
    import NamedCard from "../common/NamedCard.svelte";
    import { type Run } from "$lib/entity/api";
    let { run }: { run: Run } = $props();
</script>

<NamedCard 
    title={run.name} 
    description={run.description}
>
    <div class="d-flex justify-content-between align-items-center">
        <div>
            <p class="mb-0">Status: {run.status}</p>
            <p class="mb-0">Started At: {run.startedAt.toLocaleString()}</p>
            <p class="mb-0">Ended At: {run.endedAt.toLocaleString()}</p>
        </div>
        <span class="badge text-bg-success">
            Finished
        </span>
        <div class="ms-3">
            {#each run.parameters as parameter}
                <strong>{parameter.name}</strong>
                <div>{parameter.value}</div>
            {/each}
        </div>
        <div class="ms-3">
            {#each run.metrics as metric}
                <strong>{metric.name}</strong>
            {/each}
        </div>
        {#if run.failureReason}
            <div class="text-danger">
                <p class="mb-0">Failure Reason: {run.failureReason}</p>
            </div>
        {/if}
    </div>
    
</NamedCard>