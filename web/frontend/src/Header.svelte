<script>
  import {
    Icon,
    Collapse,
    Navbar,
    NavbarBrand,
    Nav,
    NavbarToggler,
    Dropdown,
    DropdownToggle,
    DropdownMenu,
  } from "sveltestrap";
  import NavbarLinks from "./NavbarLinks.svelte";
  import NavbarTools from "./NavbarTools.svelte";
  import NavBadge from "./navbar/nav-badge.svelte";

  export let username; // empty string if auth. is disabled, otherwise the username as string
  export let authlevel; // Integer
  export let clusters; // array of names
  export let roles; // Role Enum-Like

  let isOpen = false;
  let screenSize;

  const jobsTitle = new Map();
  jobsTitle.set(2, "Job Search");
  jobsTitle.set(3, "Managed Jobs");
  jobsTitle.set(4, "Jobs");
  jobsTitle.set(5, "Jobs");
  const usersTitle = new Map();
  usersTitle.set(3, "Managed Users");
  usersTitle.set(4, "Users");
  usersTitle.set(5, "Users");

  const views = [
    // {
    //   title: "My Jobs changed",
    //   requiredRole: roles.user,
    //   href: `/monitoring/user/${username}`,
    //   icon: "bar-chart-line-fill",
    //   perCluster: false,
    //   menu: "none",
    // },
    // {
    //   title: jobsTitle.get(authlevel),
    //   requiredRole: roles.user,
    //   href: `/monitoring/jobs/`,
    //   icon: "card-list",
    //   perCluster: false,
    //   menu: "none",
    // },
    // {
    //   title: usersTitle.get(authlevel),
    //   requiredRole: roles.manager,
    //   href: "/monitoring/users/",
    //   icon: "people-fill",
    //   perCluster: false,
    //   menu: "Groups",
    // },
    // {
    //   title: "Projects changed",
    //   requiredRole: roles.support,
    //   href: "/monitoring/projects/",
    //   icon: "folder",
    //   perCluster: false,
    //   menu: "Groups",
    // },
    // {
    //   title: "Tags",
    //   requiredRole: roles.user,
    //   href: "/monitoring/tags/",
    //   icon: "tags",
    //   perCluster: false,
    //   menu: "Groups",
    // },
    // {
    //   title: "Analysis",
    //   requiredRole: roles.support,
    //   href: "/monitoring/analysis/",
    //   icon: "graph-up",
    //   perCluster: true,
    //   menu: "Stats",
    // },
    // {
    //   title: "Nodes",
    //   requiredRole: roles.admin,
    //   href: "/monitoring/systems/",
    //   icon: "cpu",
    //   perCluster: true,
    //   menu: "Groups",
    // },
    {
      title: "Home",
      requiredRole: roles.admin,
      href: "/",
      //   icontype : "lucide",
      icon: "house",
      perCluster: false,
      menu: "menu",
    },
    {
      title: "Control",
      requiredRole: roles.admin,
      href: "/monitoring/control/",
      //   icontype : "lucide",
      icon: "folder-symlink-fill",
      perCluster: true,
      menu: "menu",
    },
    {
      title: "Partitions",
      requiredRole: roles.admin,
      href: "/monitoring/partition/",
      icon: "device-hdd",
      perCluster: true,
      menu: "menu",
    },
    {
      title: "History",
      requiredRole: roles.admin,
      href: "/monitoring/history/",
      icon: "clock-history",
      perCluster: false,
      menu: "menu",
    },
  ];
</script>

<svelte:window bind:innerWidth={screenSize} />
<Navbar color="light" light expand="md" fixed="top" style="list-style-type=none;  ">
  <NavbarBrand href="/">
    <img alt="ClusterCockpit Logo" src="/img/logo.png" height="25rem" />
  </NavbarBrand>

  {#if screenSize < 750} 
    <NavBadge />
  {/if}

  <NavbarToggler on:click={() => (isOpen = !isOpen)} />


  <Collapse
    style="justify-content: space-between"
    {isOpen}
    navbar
    expand="md"
    on:update={({ detail }) => (isOpen = detail.isOpen)}
  >
    <Nav navbar>
      {#if screenSize > 1500 || screenSize < 768}
        <NavbarLinks
          {clusters}
          links={views.filter((item) => item.requiredRole <= authlevel)}
        />
      {:else if screenSize > 1300}
        <NavbarLinks
          {clusters}
          links={views.filter(
            (item) => item.requiredRole <= authlevel && item.menu != "menu",
          )}
        />

        <Dropdown nav>
          <DropdownToggle nav caret>
            <Icon name="menu-button-wide-fill" />
            Menu
          </DropdownToggle>
          <DropdownMenu class="dropdown-menu-lg-end">
            <NavbarLinks
              {clusters}
              links={views.filter(
                (item) =>
                  item.requiredRole <= authlevel && item.menu == "menu",
              )}
            />
          </DropdownMenu>
        </Dropdown>
      {:else}
        <NavbarLinks
          {clusters}
          links={views.filter(
            (item) => item.requiredRole <= authlevel && item.menu == "none",
          )}
        />

        {#each Array("menu") as menu}
          <Dropdown nav>
            <DropdownToggle nav caret>
              {menu}
            </DropdownToggle>
            <DropdownMenu class="dropdown-menu-lg-end">
              <NavbarLinks
                {clusters}
                links={views.filter(
                  (item) => item.requiredRole <= authlevel && item.menu == menu,
                )}
              />
            </DropdownMenu>
          </Dropdown>
        {/each}
      {/if}
    </Nav>

    <NavbarTools {username} {authlevel} {roles} {screenSize} />
  </Collapse>
</Navbar>
