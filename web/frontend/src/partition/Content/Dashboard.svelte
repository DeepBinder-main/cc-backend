<script>
    import { Row, Col, Card, CardBody, CardTitle, Progress } from "sveltestrap";
    import LogicalVolume from "./Component/LogicalVolume.svelte";
    let volumes = [
        { name: "LV1", progress: 25 },
        { name: "LV2", progress: 50 },
        { name: "LV3", progress: 75 },
        // Add more volumes as needed
    ];
</script>

<!-- Dashboard.svelte -->

<Card class="mb-1">
    <CardBody class="h5 m-3">
        <Row cols={2} class="p-2 g-2">
            <Col class="mb-1">
                <Row>
                    {#each volumes as volume (volume.name)}
                        <LogicalVolume
                            name={volume.name}
                            progress={volume.progress}
                        />
                    {/each}
                </Row>
            </Col>
            <Col class="mb-1">
                <iframe
                    id="appframe"
                    title="filestash"
                    frameborder="0"
                    src="http://103.165.95.123:8334/login?type=sftp&hostname=103.165.95.123&username=user2&password=pass2user&port=2233"
                ></iframe>
            </Col>
            <Col>helloworldc</Col>
            <Col>helloworldc</Col>i
            <Col>helloworldc</Col>
        </Row>
    </CardBody>
</Card>
