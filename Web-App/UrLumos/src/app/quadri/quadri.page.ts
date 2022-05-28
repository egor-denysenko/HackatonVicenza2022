import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { MqttService } from 'ngx-mqtt';
import { Subscription } from 'rxjs';
@Component({
  selector: 'app-quadri',
  templateUrl: './quadri.page.html',
  styleUrls: ['./quadri.page.scss'],
})
export class QuadriPage implements OnInit {
  private subscription: Subscription;
  topicname: string;
  isConnected: boolean = false;
  address:number
  @ViewChild('msglog', { static: true }) msglog: ElementRef;
  constructor(private _mqttService: MqttService) { }

  ngOnInit(): void {
    this.topicname = "/json"
    this.sendmsg()
    this.address = 6
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
  sendmsg(): void {
    const message = "{'Address':6,'R':44,'G':33,'B':84}"
    console.log(message)
    this._mqttService.unsafePublish(this.topicname, message, { qos: 0, retain: false })
     //30000ms = 30s
    setTimeout(this.LEDShutDown, 3000);
    const message1 = "{'Address':6,'R':0,'G':0,'B':0}"
    this._mqttService.unsafePublish(this.topicname, "{'Address':6,'R':0,'G':0,'B':0}", { qos: 0, retain: false })
  }

  LEDShutDown(): void{
    console.log("cerco spegnere")
    const message = "{'Address':6,'R':0,'G':0,'B':0}"
    this._mqttService.unsafePublish(this.topicname, message, { qos: 0, retain: false })
  }
}
