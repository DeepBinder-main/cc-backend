<script>
    import { ModalFooter, Button, Modal, Table, Input, Container, Row, Col, Card } from "sveltestrap";
    import Agent from "./agents/agent.svelte";
    let open = false;
    const toggle = () => (open = !open);
</script>

<Container sm>
    <Row>
        <Col>
            <Card class="p-3">
                <h3>Status</h3>
                <Input name="type" type="radio" label="Active" />
                <Input name="type" type="radio" label="Disconnected" />
                <Input name="type" type="radio" label="Pending" />
            </Card>
        </Col>

        <Col>
            <Card class="p-3">
                <h3>Job Details</h3>
                <table>
                    <tr>
                        <td>Completed</td>
                        <td>00</td>
                    </tr>
                    <tr>
                        <td>Failed</td>
                        <td>00</td>
                    </tr>
                    <tr>
                        <td>Pending</td>
                        <td>00</td>
                    </tr>
                </table>
            </Card>
        </Col>
    </Row>

    <Row class="mt-5">
        <div class="d-flex  justify-content-between">
            <h2>Agents (3)</h2>
            <div>
                <Button color="primary" on:click={toggle}>Deploy New Agent</Button>
                <Button color="secondary">Refresh</Button>
            </div>
        </div>
        <Table hover>
            <thead>
                <tr>
                    <th>Id</th>
                    <th>Name</th>
                    <th>Ip Address</th>
                    <th>OS</th>
                    <th>Cluster</th>
                    <th>Status</th>
                    <th>Disconnect</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th>1</th>
                    <th>Follower</th>
                    <th>141.250.136.66</th>
                    <th>Linux</th>
                    <th>dlqMlIb</th>
                    <th>Active</th>
                    <th>
                        <Button color="danger">Disconnect</Button>
                    </th>
                </tr>
                <tr>
                    <th>2</th>
                    <th>Follower</th>
                    <th>76.220.94.207</th>
                    <th>Linux</th>
                    <th>XKc9lLn</th>
                    <th>Active</th>
                    <th>
                        <Button color="danger">Disconnect</Button>
                    </th>
                </tr>
                <tr>
                    <th>3</th>
                    <th>Leader</th>
                    <th>240.73.6.45</th>
                    <th>Linux</th>
                    <th>8LY6QaCi</th>
                    <th>Active</th>
                    <th>
                        <Button color="danger">Disconnect</Button>
                    </th>
                </tr>
            </tbody>
        </Table>
    </Row>
</Container>

<Modal body header="Deploy New Agent" isOpen={open} {toggle} size="xl">
    <Agent />
    <ModalFooter>
        <Button color="primary" on:click={toggle}>CLOSE</Button>
    </ModalFooter>
</Modal>