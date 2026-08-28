<script lang="ts">
    import type { RunDetails } from "$lib/types/base/run";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        run,
        active = false,
    }: { 
        run: RunDetails;
        active?: boolean;
    } = $props();
</script>

<MLforgeSideBarItem path={`/${run.runSummary.id}`} active={active}>
    (ID: {run.runSummary.id}) {run.runSummary.name}
</MLforgeSideBarItem>