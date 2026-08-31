<script lang="ts">
    import type { ExperimentSummary } from "$lib/types/base/experiment";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        experiment,
    }: { 
        experiment: ExperimentSummary;
    } = $props();
</script>

<MLforgeSideBarItem>
    <a href={`/browser/experiments/${experiment.id}`} class="nav-link text-white" aria-current="page"> 
        (ID: {experiment.id}) {experiment.name}
    </a>
</MLforgeSideBarItem>