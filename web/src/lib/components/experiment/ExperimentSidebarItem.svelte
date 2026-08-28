<script lang="ts">
    import type { ExperimentDetails } from "$lib/types/base/experiment";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        experiment,
        active = false,
    }: { 
        experiment: ExperimentDetails;
        active?: boolean;
    } = $props();
</script>

<MLforgeSideBarItem path={`/${experiment.experimentSummary.id}`} active={active}>
    (ID: {experiment.experimentSummary.id}) {experiment.experimentSummary.name}
</MLforgeSideBarItem>