<script lang="ts">
    import type { ExperimentSummary } from "$lib/types/base/experiment";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        experiment,
        active = false,
    }: { 
        experiment: ExperimentSummary;
        active?: boolean;
    } = $props();
</script>

<MLforgeSideBarItem path={`/browser/experiments/${experiment.id}`} active={active}>
    (ID: {experiment.id}) {experiment.name}
</MLforgeSideBarItem>