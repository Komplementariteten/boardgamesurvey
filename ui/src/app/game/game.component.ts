import { Component, Input, OnInit } from '@angular/core';
import { BoardGame } from '../boardgame';
import { ListService } from '../list.service';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {

  @Input() game: BoardGame = new BoardGame();

  selected: string = "";
  constructor(private list: ListService) { }

  ngOnInit(): void {
    this.selected = this.game.Name!;
  }
  nameChanged(ev: Event) {
    this.game.Name = this.selected;
  }
  add() {
    if (this.game != undefined){
      console.log("Add To survey");
      console.log(this.game);
      this.list.addGame(this.game);
    }
  }

}
