<script lang="ts">
    import type { Snippet } from "svelte";

    let {
        path,
        active = false,
        children
    }: { 
        path: string;
        active?: boolean;
        children: Snippet;
    } = $props();
</script>

<li class="nav-item"> 
    <a href="{path}" class="nav-link {active ? 'active' : 'text-white'}" aria-current="page"> 
        <svg class="bi pe-none me-2" width="16" height="16" aria-hidden="true">
            <use xlink:href="#home"></use>
        </svg>
        {@render children()}
    </a> 
</li>