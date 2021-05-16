import { Component, Input, OnInit } from '@angular/core';
import { ListService } from '../list.service';
import { SurveyItem } from '../surveyitem';

@Component({
  selector: 'app-voteitem',
  templateUrl: './voteitem.component.html',
  styleUrls: ['./voteitem.component.css']
})
export class VoteitemComponent implements OnInit {

  @Input() item: SurveyItem | undefined;

  showInfo: boolean;
  constructor(public votes: ListService) {
    this.showInfo = false;
  }

  vote(){
    if(this.item != undefined) {
      this.votes.vote(this.item);
    }
  }

  ngOnInit(): void {
  }

}
