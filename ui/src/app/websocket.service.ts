import { EventEmitter, Injectable } from '@angular/core';
import { Vote } from './surveyitem';


@Injectable({
  providedIn: 'root'
})
export class WebsocketService {
  private socket$: WebSocket | undefined;
  private listener: EventEmitter<any> = new EventEmitter();

  constructor() { }

  connect() {
    if (this.socket$ == undefined || this.socket$.CLOSED) {
      this.socket$ = new WebSocket("ws://" + window.location.hostname + ":" + window.location.port + "/ws");
      this.socket$.onopen = event => {
        this.listener.emit({"type": "open", "data": event});
        console.log(event);
      }
      this.socket$.onclose = event => {
        this.listener.emit({"type": "close", "data": event});
        console.log(event);
      }
      this.socket$.onmessage = event => {
        this.listener.emit({"type": "message", "data": JSON.parse(event.data)});
      }
    }
  }

  public send(msg: Vote) {
    const stringValue: string = JSON.stringify(msg);
    this.socket$?.send(stringValue);
  }
  public close(){
    this.socket$?.close();
  }
  public getEventListener() {
    return this.listener;
  }
}
