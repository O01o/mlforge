<script lang="ts">
    import Card from "$lib/components/common/Card.svelte";
    import MLforgeSideBar from "$lib/components/common/MLforgeSideBar.svelte";
    import MLforgeSideBarItem from "$lib/components/common/MLforgeSideBarItem.svelte";
    import Title from "$lib/components/common/Title.svelte";
    import NamedCard from "$lib/components/common/NamedCard.svelte";
    import type { Experiment } from "$lib/entity/api";
    import ExperimentSidebarItem from "$lib/components/experiment/ExperimentSidebarItem.svelte";
    
    const experiments = Array.from({ length: 30 }, (_, i) => ({
        id: i + 1,
        name: `Experiment ${i + 1}`,
        description: "Sample MLforge experiment",
    } as Experiment));
</script>

<div class="d-flex vh-100 overflow-hidden">

    <MLforgeSideBar>
        {#each experiments as experiment}
            <ExperimentSidebarItem experiment={experiment} active={experiment.id === 1} />
        {/each}
    </MLforgeSideBar>

    <main class="flex-grow-1 overflow-auto p-4">
        <div class="container-fluid">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <Title 
                    title="Create New Experiment" 
                    description="Create a new experiment to start tracking your training runs."
                />
                <button class="btn btn-primary">Create</button>
            </div>
            <div class="row g-3">
                <NamedCard
                    title="Experiment Name"
                    description="Experiment Name"
                >
                    <input 
                        type="text" 
                        class="form-control" 
                        id="experimentFormControlInput" 
                        placeholder="Experiment 1"
                    />
                </NamedCard>
                <NamedCard
                    title="Experiment Description"
                    description="Experiment Description"
                >
                    <textarea 
                        class="form-control" 
                        id="experimentFormControlTextarea" 
                        rows="3"
                    ></textarea>
                </NamedCard>
            </div>
        </div>
    </main>
</div>
