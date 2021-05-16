import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-infotext',
  templateUrl: './infotext.component.html',
  styleUrls: ['./infotext.component.css']
})
export class InfotextComponent implements OnInit {

  showInfoText: boolean = true;
  constructor() { }

  ngOnInit(): void {
  }

}
