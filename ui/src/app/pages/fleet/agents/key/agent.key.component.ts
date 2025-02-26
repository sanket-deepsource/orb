import { Component, Input, OnInit } from '@angular/core';
import { NbDialogRef } from '@nebular/theme';
import { STRINGS } from 'assets/text/strings';
import { ActivatedRoute, Router } from '@angular/router';
import { Agent } from 'app/common/interfaces/orb/agent.interface';

@Component({
  selector: 'ngx-agent-key-component',
  templateUrl: './agent.key.component.html',
  styleUrls: ['./agent.key.component.scss'],
})
export class AgentKeyComponent implements OnInit {
  strings = STRINGS.agents;

  command2copy: string;
  key2copy: string;

  @Input() agent: Agent = {};

  constructor(
    protected dialogRef: NbDialogRef<AgentKeyComponent>,
    protected route: ActivatedRoute,
    protected router: Router,
  ) {
  }
  ngOnInit(): void {
    this.makeCommand2Copy();
    this.key2copy = this.agent.key;
  }

  makeCommand2Copy() {
    this.command2copy = `docker run --rm --net=host\\
      \r-e ORB_CLOUD_ADDRESS=${document.location.protocol}://${document.location.hostname}/\\
      \r-e ORB_CLOUD_MQTT_ID='${this.agent.id}'\\
      \r-e ORB_CLOUD_MQTT_CHANNEL_ID='${this.agent.channel_id}'\\
      \r-e ORB_BACKENDS_PKTVISOR_IFACE=[ETH-INTERFACE]\\
      \r-e ORB_CLOUD_MQTT_KEY='${this.agent.key}''\\
      \rns1labs/orb-agent`;
  }

  onClose() {
    this.dialogRef.close(false);
  }
}
