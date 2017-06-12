package pl.cc;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Controller {

    @Autowired
    private DashboardHelloService dashboardHelloService;

    @RequestMapping("/hello")
    public String printMessage(@RequestParam(defaultValue = "pozdro600") String name) {
        return dashboardHelloService.sendMessage(name);
    }
}
