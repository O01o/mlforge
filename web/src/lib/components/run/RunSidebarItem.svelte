<script lang="ts">
    import type { Run } from "$lib/entity/api";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        run,
        active = false,
    }: { 
        run: Run;
        active?: boolean;
    } = $props();
</script>

<MLforgeSideBarItem path={`/${run.id}`} active={active}>
    (ID: {run.id}) {run.name}
</MLforgeSideBarItem>