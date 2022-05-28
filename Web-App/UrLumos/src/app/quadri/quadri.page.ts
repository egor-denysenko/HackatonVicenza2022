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
  @ViewChild('msglog', { static: true }) msglog: ElementRef;
  constructor(private _mqttService: MqttService) { }

  ngOnInit(): void {
    this.topicname = "/json"
    this.sendmsg()
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
  sendmsg(): void {
    // use unsafe publish for non-ssl websockets
    //this._mqttService.publish(this.topicname, "{'Address':1,'R':0,'G':34,'B':255}", { qos: 0, retain: false })
    this._mqttService.unsafePublish(this.topicname, "{'Address':2,'R':0,'G':34,'B':255}", { qos: 0, retain: false })
    //this.msg = ''
  }

}
