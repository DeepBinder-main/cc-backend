<script>
    import { Button, Icon, Card, Badge, CardHeader, Accordion, AccordionItem } from "sveltestrap";
</script>

<Card>
    <CardHeader>
        <Accordion flush>
            <AccordionItem>
                <div slot="header">
                    <span class="fs-3 text-bold">Notifications</span>
                    <!-- <Button color="primary" class="position-relative">
                        <Icon name="bell-fill" /> <Badge color="danger" pill positioned ariaLabel="Unread messages">3</Badge>
                    </Button> -->
                </div>
                <div>
                    <table>
                        <tr>
                            <td><Badge color="success">success</Badge> </td>
                            <td>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Doloribus, natus?</td>
                        </tr>
                    </table>
                </div>
            </AccordionItem>
        </Accordion>
    </CardHeader>
</Card>

