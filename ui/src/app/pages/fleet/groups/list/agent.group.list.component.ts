import { AfterViewChecked, AfterViewInit, ChangeDetectorRef, Component, OnInit, TemplateRef, ViewChild } from '@angular/core';
import { NbDialogService } from '@nebular/theme';

import { DropdownFilterItem } from 'app/common/interfaces/mainflux.interface';
import { ActivatedRoute, Router } from '@angular/router';
import { STRINGS } from 'assets/text/strings';
import { AgentGroupDeleteComponent } from 'app/pages/fleet/groups/delete/agent.group.delete.component';
import { AgentGroupDetailsComponent } from 'app/pages/fleet/groups/details/agent.group.details.component';
import { ColumnMode, DatatableComponent, TableColumn } from '@swimlane/ngx-datatable';
import { AgentGroupsService } from 'app/common/services/agents/agent.groups.service';
import { NgxDatabalePageInfo, OrbPagination } from 'app/common/interfaces/orb/pagination.interface';
import { AgentGroup } from 'app/common/interfaces/orb/agent.group.interface';
import { Debounce } from 'app/shared/decorators/utils';
import { AgentMatchComponent } from 'app/pages/fleet/agents/match/agent.match.component';
import { NotificationsService } from 'app/common/services/notifications/notifications.service';


@Component({
  selector: 'ngx-agent-group-list-component',
  templateUrl: './agent.group.list.component.html',
  styleUrls: ['./agent.group.list.component.scss'],
})
export class AgentGroupListComponent implements OnInit, AfterViewInit, AfterViewChecked {
  strings = STRINGS.agentGroups;

  columnMode = ColumnMode;
  columns: TableColumn[];

  loading = false;

  paginationControls: OrbPagination<AgentGroup>;

  searchPlaceholder = 'Search by name';
  filterSelectedIndex = '0';

  // templates
  @ViewChild('agentGroupTemplateCell') agentGroupsTemplateCell: TemplateRef<any>;
  @ViewChild('agentGroupTagsTemplateCell') agentGroupTagsTemplateCell: TemplateRef<any>;
  @ViewChild('actionsTemplateCell') actionsTemplateCell: TemplateRef<any>;

  tableFilters: DropdownFilterItem[] = [
    {
      id: '0',
      label: 'Name',
      prop: 'name',
      selected: false,
    },
    {
      id: '1',
      label: 'Tags',
      prop: 'tags',
      selected: false,
    },
  ];

  constructor(
    private cdr: ChangeDetectorRef,
    private dialogService: NbDialogService,
    private agentGroupsService: AgentGroupsService,
    private notificationsService: NotificationsService,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    this.agentGroupsService.clean();
    this.paginationControls = AgentGroupsService.getDefaultPagination();
  }

  @ViewChild('tableWrapper') tableWrapper;
  @ViewChild(DatatableComponent) table: DatatableComponent;
  private currentComponentWidth;
  ngAfterViewChecked() {
    if (this.table && this.table.recalculate && (this.tableWrapper.nativeElement.clientWidth !== this.currentComponentWidth)) {
      this.currentComponentWidth = this.tableWrapper.nativeElement.clientWidth;
      this.table.recalculate();
      this.cdr.detectChanges();
      window.dispatchEvent(new Event('resize'));
    }
  }

  ngOnInit() {
    this.agentGroupsService.clean();
    this.getAgentGroups();
  }

  ngAfterViewInit() {
    this.columns = [
      {
        prop: 'name',
        name: 'Name',
        resizeable: false,
        flexGrow: 1,
        minWidth: 90,
      },
      {
        prop: 'description',
        name: 'Description',
        resizeable: false,
        minWidth: 100,
        flexGrow: 2,
      },
      {
        prop: 'matching_agents',
        name: 'Agents',
        resizeable: false,
        minWidth: 100,
        flexGrow: 1,
        cellTemplate: this.agentGroupsTemplateCell,
      },
      {
        prop: 'tags',
        name: 'Tags',
        minWidth: 90,
        flexGrow: 3,
        cellClass: Object,
        cellTemplate: this.agentGroupTagsTemplateCell,
      },
      {
        name: '',
        prop: 'actions',
        minWidth: 130,
        resizeable: false,
        sortable: false,
        flexGrow: 1,
        cellTemplate: this.actionsTemplateCell,
      },
    ];

    this.cdr.detectChanges();
  }

  @Debounce(500)
  getAgentGroups(pageInfo: NgxDatabalePageInfo = null): void {
    const isFilter = this.paginationControls.name?.length > 0 || this.paginationControls.tags?.length > 0;

    if (isFilter) {
      pageInfo = {
        offset: this.paginationControls.offset,
        limit: this.paginationControls.limit,
      };
      if (this.paginationControls.name?.length > 0) pageInfo.name = this.paginationControls.name;
      if (this.paginationControls.tags?.length > 0) pageInfo.tags = this.paginationControls.tags;
    }

    this.loading = true;
    this.agentGroupsService.getAgentGroups(pageInfo, isFilter).subscribe(
      (resp: OrbPagination<AgentGroup>) => {
        this.paginationControls = resp;
        this.paginationControls.offset = pageInfo.offset;
        this.paginationControls.total = resp.total;
        this.loading = false;
      },
    );
  }

  onOpenAdd() {
    this.router.navigate(['add'], {
      relativeTo: this.route,
    });
  }

  onOpenEdit(agentGroup: any) {
    this.router.navigate([`edit/${ agentGroup.id }`], {
      state: { agentGroup: agentGroup, edit: true },
      relativeTo: this.route,
    });
  }

  onFilterSelected(selectedIndex) {
    this.searchPlaceholder = `Search by ${ this.tableFilters[selectedIndex].label }`;
  }

  openDeleteModal(row: any) {
    const { name, id } = row;
    this.dialogService.open(AgentGroupDeleteComponent, {
      context: { name },
      autoFocus: true,
      closeOnEsc: true,
    }).onClose.subscribe(
      confirm => {
        if (confirm) {
          this.agentGroupsService.deleteAgentGroup(id).subscribe(() => {
            this.notificationsService.success('Agent Group successfully deleted', '');
            this.getAgentGroups();
          });
        }
      },
    );
  }

  openDetailsModal(row: any) {
    this.dialogService.open(AgentGroupDetailsComponent, {
      context: { agentGroup: row },
      autoFocus: true,
      closeOnEsc: true,
    }).onClose.subscribe((resp) => {
      if (resp) {
        this.onOpenEdit(row);
      } else {
        this.getAgentGroups();
      }
    });
  }

  onMatchingAgentsModal(row: any) {
    this.dialogService.open(AgentMatchComponent, {
      context: {agentGroup: row},
      autoFocus: true,
      closeOnEsc: true,
    }).onClose.subscribe(_ => {
      this.getAgentGroups();
    });
  }

  searchAgentByName(input) {
    this.getAgentGroups({
      ...this.paginationControls,
      [this.tableFilters[this.filterSelectedIndex].prop]: input,
    });
  }
}
