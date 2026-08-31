<script lang="ts">
    import type { RunDetails } from "$lib/types/base/run";
    import MLforgeSideBarItem from "../common/MLforgeSideBarItem.svelte";
    let {
        run,
        active = $bindable<boolean>(),
    }: { 
        run: RunDetails;
        active: boolean;
    } = $props();
</script>

<MLforgeSideBarItem>
    <button
        type="button"
        class="nav-link {active ? 'active' : 'text-white'}" 
        aria-current="page"
        on:click={() => {
            active = !active;
        }}
    > 
        <svg class="bi pe-none me-2" width="16" height="16" aria-hidden="true">
            <use xlink:href="#home"></use>
        </svg>
        (ID: {run.runSummary.id}) {run.runSummary.name}
    </button>
</MLforgeSideBarItem>