<script>
    import {Button} from 'sveltestrap';
</script>

<div class="row mt-3">
    <div class="col-10">
        <div class="d-flex">
            <p class="fw-bold fs-4 mt-2">Storage Maintenance </p>
            <select name="user" id="user" style="height: 50px;" class="w-75 ms-3 form-select border-dark">
                <option value="">select user</option>
                <option value="">Sam Hunt</option>
                <option value="">Victoria Haynes</option>
                <option value="">Cecilia Price</option>
            </select>
        </div>
    </div>
    <div class="col-2">
        <Button outline color="dark">Attach Partition</Button>
    </div>
</div>

<table class="table mt-3 table-hover table-bordered">
    <thead class="table-dark">
        <tr>
            <th>User</th>
            <th>Volume Group</th>
            <th>Logical Volume </th>
            <th>Created At </th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>User1</td>
            <td>vg</td>
            <td>lv</td>
            <td>10-19-2119 00:00:00</td>
        </tr>
        <tr>
            <td>User2</td>
            <td>vg</td>
            <td>lv</td>
            <td>2-27-2073 00:00:00</td>
        </tr>
        <tr>
            <td>User3</td>
            <td>vg</td>
            <td>lv</td>
            <td>11-11-2094 00:00:00</td>
        </tr>
    </tbody>
</table>
