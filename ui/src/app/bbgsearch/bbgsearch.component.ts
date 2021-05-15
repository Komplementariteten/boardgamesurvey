import { Component, OnInit } from '@angular/core';
import {BbgService} from '../bbg.service'

@Component({
  selector: 'app-bbgsearch',
  templateUrl: './bbgsearch.component.html',
  styleUrls: ['./bbgsearch.component.css']
})
export class BbgsearchComponent implements OnInit {

  constructor(private bbgService: BbgService) { }

  ngOnInit(): void {
    this.bbgService.search("hallo").subscribe(d => {
      console.log(d);
    });
  }

}
