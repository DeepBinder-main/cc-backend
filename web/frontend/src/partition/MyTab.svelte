<script>
    import { TabContent, TabPane } from "sveltestrap";
    import DashboardTab from "./tabs/Dashboard.svelte"
    import LVMManager from "./tabs/CreateLVM.svelte"
    import AttachLVM from "./tabs/AttachLVM.svelte"
    import SettingTab from "./tabs/Setting.svelte"
    import AutoStorage from "./tabs/AutoStorage.svelte";
</script>

<TabContent  class="p-3">
    <TabPane tabId="dashboard" active>
        <span class="fs-4" slot="tab">
            Dashboard
        </span>
        <DashboardTab />
    </TabPane>
    <TabPane tabId="create-lvm">
        <span class="fs-4" slot="tab">
            LVM Manager
        </span>

        <LVMManager />
    </TabPane>
    
    <TabPane tabId="monitor-tab">
        <span class="fs-4" slot="tab">
            Attach Volume
        </span>

        <AttachLVM />
    </TabPane>
    
    <TabPane tabId="auto-storage">
        <span class="fs-4" slot="tab">
            Auto Storage
        </span>

        <AutoStorage />
    </TabPane>
    
    <!-- <TabPane tabId="setting-tab">
        <span class="fs-4" slot="tab">
            Setting
        </span>

        <SettingTab />
    </TabPane> -->
    
</TabContent>
