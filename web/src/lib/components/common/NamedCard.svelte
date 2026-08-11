<script lang="ts">
    import type { Snippet } from "svelte";
    import Card from "./Card.svelte";

    let {
        title,
        description,
        children
    }: {
        title: string;
        description: string;
        children: Snippet;
    } = $props();
</script>

<Card>
    <div class="justify-content-between align-items-start">
        <div class="mb-3">
            <h2 class="h5 card-title">{title}</h2>
            <p class="card-text text-body-secondary">{description}</p>
        </div>
        {@render children()}
    </div>
</Card>