<script lang="ts">
    import type { Experiment } from "$lib/entity/api";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        experiment,
        active = false,
    }: { 
        experiment: Experiment;
        active?: boolean;
    } = $props();
</script>

<MLforgeSideBarItem path={`/${experiment.id}`} active={active}>
    (ID: {experiment.id}) {experiment.name}
</MLforgeSideBarItem>