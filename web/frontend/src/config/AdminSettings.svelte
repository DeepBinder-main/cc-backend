<script>
    import { Row, Col } from 'sveltestrap'
    import { onMount } from 'svelte'
    import EditRole from './admin/EditRole.svelte'
    import EditProject from './admin/EditProject.svelte'
    import AddUser from './admin/AddUser.svelte'
    import ShowUsers from './admin/ShowUsers.svelte'
    import Options from './admin/Options.svelte'

    let users = []
    let roles = []

    function getUserList() {
        fetch('/api/users/?via-ldap=false&not-just-user=true')
            .then(res => res.json())
            .then(usersRaw => {
                users = usersRaw
        })
    }

    function getValidRoles() {
        fetch('/api/roles/')
            .then(res => res.json())
            .then(rolesRaw => {
                roles = rolesRaw
        })
    }

    function initAdmin() {
      getUserList()
      getValidRoles()
    }

    onMount(() => initAdmin())

</script>

<Row cols={2} class="p-2 g-2" >
    <Col class="mb-1">
        <AddUser roles={roles} on:reload={getUserList}/>
    </Col>
    <Col class="mb-1">
        <ShowUsers on:reload={getUserList} bind:users={users}/>
    </Col>
    <Col>
        <EditRole roles={roles} on:reload={getUserList}/>
    </Col>
    <Col>
        <EditProject on:reload={getUserList}/>
    </Col>
    <Col>
        <Options/>
    </Col>
</Row>
