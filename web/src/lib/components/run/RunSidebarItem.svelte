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

<MLforgeSideBarItem>
    <div class="nav-link {active ? 'active' : 'text-white'}" aria-current="page"> 
        <svg class="bi pe-none me-2" width="16" height="16" aria-hidden="true">
            <use xlink:href="#home"></use>
        </svg>
        (ID: {run.runSummary.id}) {run.runSummary.name}
    </div>
</MLforgeSideBarItem>