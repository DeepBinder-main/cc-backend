<script>
    import {Button, Modal, ModalHeader, Input, ModalBody} from 'sveltestrap';
    import AddUser from './AddUser.svelte';
    let open = false;
    const toggle = () => (open = !open);
</script>

<div class="d-flex gap-3">
    <div class="col-2">
        <div class="d-grid" >
            <Button on:click={toggle}>Add Users</Button>
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
            <th>Username</th>
            <th>Group</th>
            <th>Last Active</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>Maude</td>
            <td>Group1</td>
            <td>9/6/2051</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Harry</td>
            <td>Group1</td>
            <td>7/12/2029</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Anthony</td>
            <td>Group1</td>
            <td>7/5/2030</td>
        </tr>
    </tbody>
</table>

<Modal body header="Add Or Edit User" isOpen={open} {toggle} size="xl">
    <AddUser />
</Modal>
