<script>
    import {Button, Modal, ModalHeader, Input, ModalBody} from 'sveltestrap';
    import AddGroup from './AddGroup.svelte';
    let open = false;
    const toggle = () => (open = !open);
</script>

<div class="d-flex gap-3">
    <div class="col-2">
        <div class="d-grid" >
            <Button on:click={toggle}>Add Group</Button>
        </div>
    </div>
    <div class="col">
        <Input placeholder="search..." />
    </div>
</div>
<table class="table mt-3 table-bordered table-hover">
    <thead class="table-dark">
        <tr>
            <th>Id</th>
            <th>Name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>Group1</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Group2</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Group3</td>
        </tr>
    </tbody>
</table>

<Modal body header="Add Or Edit Group" isOpen={open} {toggle} size="xl">
    <AddGroup />
</Modal>
