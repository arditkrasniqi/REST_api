<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\User;

class UsersController extends Controller {

    const MODEL = "App\User";

    public function all(){
        return User::all();
    }

    public function getFirst(){
        return User::where('id','>',5)->get();
    }

}
