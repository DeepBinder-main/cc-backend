<script>
    import { Icon } from "sveltestrap";
    // import CircleCheckIcon from "./CircleCheckIcon.svelte";
    // import LaptopIcon from "./LaptopIcon.svelte";
    // import AppWindowIcon from "./AppWindowIcon.svelte";
    // import AppleIcon from "./AppleIcon.svelte";
</script>

<div class="flex flex-col bg-white p-6 rounded-lg shadow-md w-full max-w-2xl">
    <div class="flex items-start justify-between">
        <h2 class="text-2xl font-semibold">Deploy new agent</h2>
        <button
            class="inline-flex items-center justify-center whitespace-nowrap rounded-md font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 text-sm"
        >
            Close
        </button>
    </div>
    <div class="mt-4">
        <div class="flex items-center">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="text-green-500 mr-2"
            >
                <circle cx="12" cy="12" r="10"></circle>
                <path d="m9 12 2 2 4-4"></path>
            </svg>
            <h3 class="text-lg font-medium">
                Select the package to download and install on your system:
            </h3>
        </div>
        <div class="mt-2 grid grid-cols-3 gap-4">
            <div class="col-span-2 grid grid-cols-2 gap-4">
                <div class="flex items-center">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-gray-500 mr-2"
                    >
                        <path
                            d="M20 16V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v9m16 0H4m16 0 1.28 2.55a1 1 0 0 1-.9 1.45H3.62a1 1 0 0 1-.9-1.45L4 16"
                        ></path>
                    </svg>
                    <div>
                        <div
                            role="radiogroup"
                            aria-required="false"
                            dir="ltr"
                            class="grid gap-2"
                            tabindex="0"
                            style="outline: none;"
                        >
                            <div class="flex items-center space-x-2">
                                <button
                                    type="button"
                                    role="radio"
                                    aria-checked="true"
                                    data-state="checked"
                                    value="rpm-amd64"
                                    class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="rpm-amd64"
                                    tabindex="-1"
                                    data-radix-collection-item=""
                                >
                                    <span
                                        data-state="checked"
                                        class="flex items-center justify-center"
                                    >
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="24"
                                            height="24"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="2"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            class="lucide lucide-circle h-2.5 w-2.5 fill-current text-current"
                                        >
                                            <circle cx="12" cy="12" r="10"
                                            ></circle>
                                        </svg>
                                    </span>
                                </button>
                                <label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="rpm-amd64"
                                >
                                    RPM amd64
                                </label>
                            </div>
                            <div class="flex items-center space-x-2">
                                <button
                                    type="button"
                                    role="radio"
                                    aria-checked="false"
                                    data-state="unchecked"
                                    value="rpm-aarch64"
                                    class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="rpm-aarch64"
                                    tabindex="-1"
                                    data-radix-collection-item=""
                                ></button>
                                <label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="rpm-aarch64"
                                >
                                    RPM aarch64
                                </label>
                            </div>
                            <div class="flex items-center space-x-2">
                                <button
                                    type="button"
                                    role="radio"
                                    aria-checked="false"
                                    data-state="unchecked"
                                    value="deb-amd64"
                                    class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="deb-amd64"
                                    tabindex="-1"
                                    data-radix-collection-item=""
                                ></button>
                                <label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="deb-amd64"
                                >
                                    DEB amd64
                                </label>
                            </div>
                            <div class="flex items-center space-x-2">
                                <button
                                    type="button"
                                    role="radio"
                                    aria-checked="false"
                                    data-state="unchecked"
                                    value="deb-aarch64"
                                    class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="deb-aarch64"
                                    tabindex="-1"
                                    data-radix-collection-item=""
                                ></button>
                                <label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="deb-aarch64"
                                >
                                    DEB aarch64
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex items-center">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="text-gray-500 mr-2"
                    >
                        <rect x="2" y="4" width="20" height="16" rx="2"></rect>
                        <path d="M10 4v4"></path>
                        <path d="M2 8h20"></path>
                        <path d="M6 4v4"></path>
                    </svg>
                    <div>
                        <div
                            role="radiogroup"
                            aria-required="false"
                            dir="ltr"
                            class="grid gap-2"
                            tabindex="0"
                            style="outline: none;"
                        >
                            <div class="flex items-center space-x-2">
                                <button
                                    type="button"
                                    role="radio"
                                    aria-checked="true"
                                    data-state="checked"
                                    value="msi-32"
                                    class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="msi-32"
                                    tabindex="-1"
                                    data-radix-collection-item=""
                                >
                                    <span
                                        data-state="checked"
                                        class="flex items-center justify-center"
                                    >
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="24"
                                            height="24"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="2"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            class="lucide lucide-circle h-2.5 w-2.5 fill-current text-current"
                                        >
                                            <circle cx="12" cy="12" r="10"
                                            ></circle>
                                        </svg>
                                    </span>
                                </button>
                                <label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="msi-32"
                                >
                                    MSI 32/64 bits
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex items-center">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="text-gray-500 mr-2"
                >
                    <path
                        d="M12 20.94c1.5 0 2.75 1.06 4 1.06 3 0 6-8 6-12.22A4.91 4.91 0 0 0 17 5c-2.22 0-4 1.44-5 2-1-.56-2.78-2-5-2a4.9 4.9 0 0 0-5 4.78C2 14 5 22 8 22c1.25 0 2.5-1.06 4-1.06Z"
                    ></path>
                    <path d="M10 2c1 .5 2 2 2 5"></path>
                </svg>
                <div>
                    <div
                        role="radiogroup"
                        aria-required="false"
                        dir="ltr"
                        class="grid gap-2"
                        tabindex="0"
                        style="outline: none;"
                    >
                        <div class="flex items-center space-x-2">
                            <button
                                type="button"
                                role="radio"
                                aria-checked="true"
                                data-state="checked"
                                value="intel"
                                class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                id="intel"
                                tabindex="-1"
                                data-radix-collection-item=""
                            >
                                <span
                                    data-state="checked"
                                    class="flex items-center justify-center"
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        width="24"
                                        height="24"
                                        viewBox="0 0 24 24"
                                        fill="none"
                                        stroke="currentColor"
                                        stroke-width="2"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        class="lucide lucide-circle h-2.5 w-2.5 fill-current text-current"
                                    >
                                        <circle cx="12" cy="12" r="10"></circle>
                                    </svg>
                                </span>
                            </button>
                            <label
                                class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                for="intel"
                            >
                                Intel
                            </label>
                        </div>
                        <div class="flex items-center space-x-2">
                            <button
                                type="button"
                                role="radio"
                                aria-checked="false"
                                data-state="unchecked"
                                value="apple-silicon"
                                class="aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                id="apple-silicon"
                                tabindex="-1"
                                data-radix-collection-item=""
                            ></button>
                            <label
                                class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                for="apple-silicon"
                            >
                                Apple silicon
                            </label>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="mt-2">
            <a class="text-blue-600" href="#" rel="ugc">
                For additional systems and architectures, please check our
                documentation.
            </a>
        </div>
    </div>
    <div class="mt-4">
        <div class="flex items-center">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="text-green-500 mr-2"
            >
                <circle cx="12" cy="12" r="10"></circle>
                <path d="m9 12 2 2 4-4"></path>
            </svg>
            <h3 class="text-lg font-medium">Server address</h3>
        </div>
        <p class="mt-2 text-sm">
            This is the address the agent uses to communicate with the Wazuh
            server. Enter an IP address or a fully qualified domain name (FQDN).
        </p>
        <div class="mt-2">
            <label
                class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 sr-only"
                for="server-address"
            >
                Assign a server address
            </label>
            <input
                class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                id="server-address"
                placeholder="10.0.0.198"
            />
        </div>
    </div>
    <div class="mt-4">
        <div class="flex items-center">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="text-green-500 mr-2"
            >
                <circle cx="12" cy="12" r="10"></circle>
                <path d="m9 12 2 2 4-4"></path>
            </svg>
            <h3 class="text-lg font-medium">Optional settings</h3>
        </div>
        <p class="mt-2 text-sm">
            The deployment sets the endpoint hostname as the agent name by
            default. Optionally, you can set your own name in the field below.
        </p>
        <div class="mt-2">
            <label
                class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 sr-only"
                for="agent-name"
            >
                Assign an agent name
            </label>
            <input
                class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                id="agent-name"
                placeholder="Windows"
            />
        </div>
        <p class="mt-2 text-sm text-yellow-600">
            The agent name must be unique. It can't be changed once the agent
            has been enrolled.
        </p>
        <div class="mt-2">
            <label
                class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 sr-only"
                for="agent-groups"
            >
                Select one or more existing groups
            </label>
        </div>
    </div>
    <div class="mt-4">
        <div class="flex items-center">
            <span class="text-blue-600 font-bold text-lg">4</span>
            <h3 class="text-lg font-medium ml-2">
                Run the following commands to download and install the Wazuh
                agent:
            </h3>
        </div>
        <pre class="mt-2 bg-gray-100 p-4 rounded-md">
        <code>
          Invoke-WebRequest -Uri https://packages.wazuh.com/4.x/windows/wazuh-agent-4.6.0-1.msi -OutFile
          $(env:tmp)\wazuh-agent; msiexec.exe /i $(env:tmp)\wazuh-agent /q WAZUH_MANAGER='10.0.0.198'
          WAZUH_AGENT_GROUP='windows' WAZUH_AGENT_NAME='Windows' WAZUH_REGISTRATION_SERVER='10.0.0.198'
        </code>
      </pre>
    </div>
    <div class="mt-4">
        <div class="flex items-center">
            <span class="text-blue-600 font-bold text-lg">5</span>
            <h3 class="text-lg font-medium ml-2">Start the Wazuh agent:</h3>
        </div>
        <pre class="mt-2 bg-gray-100 p-4 rounded-md">
        <code>NET START WazuhSvc</code>
      </pre>
    </div>
    <div class="mt-6 flex justify-end">
        <button
            class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2"
        >
            Close
        </button>
    </div>
</div>
