import { Component, OnInit, HostBinding } from '@angular/core';
import {
  trigger,
  state,
  style,
  animate,
  stagger,
  transition,
  keyframes} from '@angular/animations';
import { setClassMetadata } from '@angular/core/src/r3_symbols';
@Component({
  animations: [
    trigger('highlighHideButton', [
    state('initial', style({transform: 'scale(1)'})),
    transition('* => initial', [animate('1.8s', keyframes([
     style({transform: 'scale(1)'}),
     style({transform: 'scale(1.1)'}),
     style({transform: 'scale(1.3)'}),
     style({transform: 'scale(1.1,1.1)'}),
     style({transform: 'scale(1.3,1.3)'}),
     style({transform: 'scale(1,1)'}),
    ]))]),
    ]),
    trigger('animateArrow', [
      state('initial', style({transform: 'scale(1)', opacity: 0})),
      transition('* => initial', [animate('3s', keyframes([
       style({transform: 'scale(1)'}),
       style({transform: 'scale(1.1)', opacity: 0.9}),
       style({transform: 'scale(1.3)', opacity: 0.7}),
       style({transform: 'scale(1.1,1.1)', opacity: 0.5}),
       style({transform: 'scale(1.3,1.3)', opacity: 0.3}),
       style({transform: 'scale(1,1)', opacity: 0}),
      ]))]),
      ])
  ],
  selector: 'app-infotext',
  templateUrl: './infotext.component.html',
  styleUrls: ['./infotext.component.css']
})
export class InfotextComponent implements OnInit {

  showInfoText: boolean = true;
  bounceInfoText: string = 'initial';
  hasScrollbar: boolean = false;
  constructor() {
  }
  hideInfoText() {
    this.showInfoText = false;
    if(document.firstElementChild != null) {
      this.hasScrollbar = document.firstElementChild.clientHeight < document.firstElementChild.scrollHeight;
    }
  }
  ngOnInit(): void {
    if(document.firstElementChild != null) {
      this.hasScrollbar = document.firstElementChild.clientHeight < document.firstElementChild.scrollHeight;
    }
  }

}
