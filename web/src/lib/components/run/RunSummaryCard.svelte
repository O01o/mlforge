<script lang="ts">
    import NamedCard from "../common/NamedCard.svelte";
    import type { RunDetails } from "$lib/types/base/run";
    let { run }: { run: RunDetails } = $props();
</script>

<NamedCard 
    title={run.runSummary.name} 
    description={run.runSummary.description}
>
    <div class="d-flex justify-content-between align-items-center">
        <div>
            <p class="mb-0">Status: {run.runSummary.status}</p>
            <p class="mb-0">Started At: {run.runSummary.startedAt.toLocaleString()}</p>
            <p class="mb-0">Ended At: {run.runSummary.endedAt.toLocaleString()}</p>
        </div>
        {#if run.runSummary.status === 1}
            <span class="badge text-bg-success">
                Finished
            </span>
        {/if}
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
        {#if run.runSummary.failureReason}
            <div class="text-danger">
                <p class="mb-0">Failure Reason: {run.runSummary.failureReason}</p>
            </div>
        {/if}
    </div>
    
</NamedCard>