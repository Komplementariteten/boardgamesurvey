import { Component, OnInit } from '@angular/core';
import { ListService } from '../list.service';
import { SurveyItem } from '../surveyitem';

@Component({
  selector: 'app-categorie',
  templateUrl: './categorie.component.html',
  styleUrls: ['./categorie.component.css']
})
export class CategorieComponent implements OnInit {

  title: string | undefined;
  description: string | undefined;
  constructor(private list: ListService) { }

  ngOnInit(): void {
  }

  add() {
    let item = new SurveyItem();
    item.Title = this.title;
    item.Description = this.description;
    this.list.add(item);
  }

}
