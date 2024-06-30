<script>
    import {Input, Container, Button} from 'sveltestrap'
</script>

<Container>
    
    <h1>Edit Group</h1>
    
    <div class="d-flex mt-3">
        <label class="col-1" for="name">Group Name : </label>
        <div class="col-11" type="text" >
            <input type="text" class="form-control" id="">
        </div>
    </div>
    <hr >
    
    <div class="mt-3">
        <h4>Group Restrictions</h4>
        <Input type="checkbox" label="Disable" />
    </div>
    <hr >
    
    <div class="mt-3">
        <h4>Permission</h4>
        <Input type="checkbox" label="Admin" />
        <Input type="checkbox" label="Create new User" />
        <Input type="checkbox" label="Create new User Group" />
    </div>
    <hr>
    
    <div class="mt-3">
        <h4>Users</h4>
        <Input type="checkbox" label="User1" />
        <Input type="checkbox" label="User2" />
        <Input type="checkbox" label="User3" />
    </div>
    <hr>

</Container>
