<script lang="ts">
    import type { Snippet } from "svelte";

    let {
        children
    }: {
        children: Snippet;
    } = $props();
</script>
    <div class="col-12">
        <div class="card">
            <div class="card-body">
                {@render children()}
            </div>
        </div>
    </div>